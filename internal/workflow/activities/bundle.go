package activities

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	securejoin "github.com/cyphar/filepath-securejoin"
	"github.com/mholt/archiver/v3"
	"github.com/otiai10/copy"
	temporal_tools "go.artefactual.dev/tools/temporal"

	"github.com/artefactual-sdps/enduro/internal/bundler"
	"github.com/artefactual-sdps/enduro/internal/watcher"
)

type BundleActivity struct {
	wsvc watcher.Service
}

func NewBundleActivity(wsvc watcher.Service) *BundleActivity {
	return &BundleActivity{wsvc: wsvc}
}

type BundleActivityParams struct {
	WatcherName      string
	TransferDir      string
	Key              string
	TempFile         string
	StripTopLevelDir bool
	IsDir            bool
	// Will override certain preservation actions that are unncesary due to the SFA-preprocessing step.
	OverrideWatcher bool
}

type BundleActivityResult struct {
	RelPath             string // Path of the transfer relative to the transfer directory.
	FullPath            string // Full path to the transfer in the worker running the session.
	FullPathBeforeStrip string // Same as FullPath but includes the top-level dir even when stripped.
}

func (a *BundleActivity) Execute(ctx context.Context, params *BundleActivityParams) (*BundleActivityResult, error) {
	var (
		res = &BundleActivityResult{}
		err error
	)

	if params.TransferDir == "" {
		params.TransferDir, err = os.MkdirTemp("", "*-enduro-transfer")
		if err != nil {
			return nil, err
		}
	}

	defer func() {
		if err != nil {
			err = temporal_tools.NewNonRetryableError(err)
		}
	}()

	if params.IsDir {
		if params.OverrideWatcher {
			src := params.TempFile
			dst := params.TransferDir
			res.FullPath, res.FullPathBeforeStrip, err = a.Copy(ctx, src, dst, false)
		} else {
			var w watcher.Watcher
			w, err = a.wsvc.ByName(params.WatcherName)
			if err == nil {
				src, _ := securejoin.SecureJoin(w.Path(), params.Key)
				dst := params.TransferDir
				res.FullPath, res.FullPathBeforeStrip, err = a.Copy(ctx, src, dst, false)
			}
		}
	} else {
		unar := a.Unarchiver(params.Key, params.TempFile)
		if unar == nil {
			res.FullPath, err = a.SingleFile(ctx, params.TransferDir, params.Key, params.TempFile)
			res.FullPathBeforeStrip = res.FullPath
		} else {
			res.FullPath, res.FullPathBeforeStrip, err = a.Bundle(ctx, unar, params.TransferDir, params.Key, params.TempFile, params.StripTopLevelDir)
		}
	}
	if err != nil {
		return nil, temporal_tools.NewNonRetryableError(err)
	}

	err = unbag(res.FullPath)
	if err != nil {
		return nil, temporal_tools.NewNonRetryableError(err)
	}

	res.RelPath, err = filepath.Rel(params.TransferDir, res.FullPath)
	if err != nil {
		return nil, fmt.Errorf("error calculating relative path to transfer (base=%q, target=%q): %v", params.TransferDir, res.FullPath, err)
	}

	return res, err
}

// Unarchiver returns the unarchiver suited for the archival format.
func (a *BundleActivity) Unarchiver(key, filename string) archiver.Unarchiver {
	if iface, err := archiver.ByExtension(key); err == nil {
		if u, ok := iface.(archiver.Unarchiver); ok {
			return u
		}
	}

	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil
	}
	defer file.Close() //#nosec G307 -- Errors returned by Close() here do not require specific handling.
	if u, err := archiver.ByHeader(file); err == nil {
		return u
	}

	return nil
}

// SingleFile bundles a transfer with the downloaded blob in it.
func (a *BundleActivity) SingleFile(ctx context.Context, transferDir, key, tempFile string) (string, error) {
	b, err := bundler.NewBundlerWithTempDir(transferDir)
	if err != nil {
		return "", fmt.Errorf("error creating bundle: %v", err)
	}

	dest, err := b.Create(filepath.Join("objects", key))
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dest.Close()

	path, _ := securejoin.SecureJoin(transferDir, dest.Name())
	if err := os.Rename(tempFile, path); err != nil {
		return "", fmt.Errorf("error moving file (from %s to %s): %v", tempFile, path, err)
	}

	if err := os.Chmod(path, os.FileMode(0o755)); err != nil {
		return "", fmt.Errorf("error changing file mode: %v", err)
	}

	if err := b.Bundle(); err != nil {
		return "", fmt.Errorf("error bundling the transfer: %v", err)
	}

	return b.FullBaseFsPath(), nil
}

// Bundle a transfer with the contents found in the archive.
func (a *BundleActivity) Bundle(ctx context.Context, unar archiver.Unarchiver, transferDir, key, tempFile string, stripTopLevelDir bool) (string, string, error) {
	// Create a new directory for our transfer with the name randomized.
	const prefix = "enduro"
	tempDir, err := os.MkdirTemp(transferDir, prefix)
	if err != nil {
		return "", "", fmt.Errorf("error creating temporary directory: %s", err)
	}
	_ = os.Chmod(tempDir, os.FileMode(0o755))

	if err := unar.Unarchive(tempFile, tempDir); err != nil {
		return "", "", fmt.Errorf("error unarchiving file: %v", err)
	}

	tempDirBeforeStrip := tempDir
	if stripTopLevelDir {
		tempDir, err = stripDirContainer(tempDir)
		if err != nil {
			return "", "", err
		}
	}

	// Delete the archive. We still have a copy in the watched source.
	_ = os.Remove(tempFile)

	return tempDir, tempDirBeforeStrip, nil
}

// Copy a transfer in the given destination using an intermediate temp. directory.
func (a *BundleActivity) Copy(ctx context.Context, src, dst string, stripTopLevelDir bool) (string, string, error) {
	const prefix = "enduro"
	tempDir, err := os.MkdirTemp(dst, prefix)
	if err != nil {
		return "", "", fmt.Errorf("error creating temporary directory: %s", err)
	}
	_ = os.Chmod(tempDir, os.FileMode(0o755))

	if err := copy.Copy(src, tempDir); err != nil {
		return "", "", fmt.Errorf("error copying transfer: %v", err)
	}

	tempDirBeforeStrip := tempDir
	if stripTopLevelDir {
		tempDir, err = stripDirContainer(tempDir)
		if err != nil {
			return "", "", err
		}
	}

	return tempDir, tempDirBeforeStrip, nil
}

// stripDirContainer strips the top-level directory of a transfer.
func stripDirContainer(path string) (string, error) {
	const errPrefix = "error stripping top-level dir"
	ff, err := os.Open(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("%s: cannot open path: %v", errPrefix, err)
	}
	fis, err := ff.Readdir(2)
	if err != nil {
		return "", fmt.Errorf("%s: error reading dir: %v", errPrefix, err)
	}
	if len(fis) != 1 {
		return "", fmt.Errorf("%s: directory has more than one child", errPrefix)
	}
	if !fis[0].IsDir() {
		return "", fmt.Errorf("%s: top-level item is not a directory", errPrefix)
	}
	securePath, _ := securejoin.SecureJoin(path, fis[0].Name())
	return securePath, nil
}

// unbag converts a bagged transfer into a standard Archivematica transfer.
// It returns a nil error if a bag is not identified, and non-nil errors when
// the bag seems invalid, without verifying the actual file contents.
func unbag(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("not a directory")
	}

	// Only continue if we have a bag.
	securePath, _ := securejoin.SecureJoin(path, "bagit.txt")
	_, err = os.Stat(securePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	// Confirm completeness of the bag.
	// if err := bagit.Valid(path); err != nil {
	// 	return err
	// }

	// Move files in data up one level if 'objects' folder already exists.
	// Otherwise, rename data to objects.
	dataPath, _ := securejoin.SecureJoin(path, "data")
	if fi, err := os.Stat(dataPath); !os.IsNotExist(err) && fi.IsDir() {
		items, err := os.ReadDir(dataPath)
		if err != nil {
			return err
		}
		for _, item := range items {
			src, _ := securejoin.SecureJoin(dataPath, item.Name())
			dst, _ := securejoin.SecureJoin(path, filepath.Base(src))
			if err := os.Rename(src, dst); err != nil {
				return err
			}
		}
		if err := os.RemoveAll(dataPath); err != nil {
			return err
		}
	} else {
		dst, _ := securejoin.SecureJoin(path, "objects")
		if err := os.Rename(dataPath, dst); err != nil {
			return err
		}
	}

	// Create metadata and submissionDocumentation directories.
	metadataPath, _ := securejoin.SecureJoin(path, "metadata")
	documentationPath, _ := securejoin.SecureJoin(metadataPath, "submissionDocumentation")
	//#nosec G301 -- Evaluate use of UID and GID among containers so that permission 750 could be used.
	err = os.MkdirAll(metadataPath, 0o775)
	if err != nil {
		return err
	}
	//#nosec G301 -- Evaluate use of UID and GID among containers so that permission 750 could be used.
	err = os.MkdirAll(documentationPath, 0o775)
	if err != nil {
		return err
	}

	// Write manifest checksums to checksum file.
	for _, item := range [][2]string{
		{"manifest-sha512.txt", "checksum.sha512"},
		{"manifest-sha256.txt", "checksum.sha256"},
		{"manifest-sha1.txt", "checksum.sha1"},
		{"manifest-md5.txt", "checksum.md5"},
	} {
		securePath, _ := securejoin.SecureJoin(path, item[0])
		file, err := os.Open(securePath) //#nosec G304 -- Potential file inclusion not possible. item[0] is coming from controlled list.
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}
		defer file.Close() //#nosec G307 -- Errors returned by Close() here do not require specific handling.

		securePath, _ = securejoin.SecureJoin(metadataPath, item[1])
		newFile, err := os.Create(securePath) //#nosec G304 -- Potential file inclusion not possible. item[1] is coming from controlled list.
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}
		defer newFile.Close() //#nosec G307 -- Errors returned by Close() here do not require specific handling.

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			newLine := ""
			if strings.Contains(line, "data/objects/") {
				newLine = strings.Replace(line, "data/objects/", "../objects/", 1)
			} else {
				newLine = strings.Replace(line, "data/", "../objects/", 1)
			}
			fmt.Fprintln(newFile, newLine)
		}

		break // One file is enough.
	}

	// Move bag files to submissionDocumentation.
	for _, item := range []string{
		"bag-info.txt",
		"bagit.txt",
		"manifest-md5.txt",
		"tagmanifest-md5.txt",
		"manifest-sha1.txt",
		"tagmanifest-sha1.txt",
		"manifest-sha256.txt",
		"tagmanifest-sha256.txt",
		"manifest-sha512.txt",
		"tagmanifest-sha512.txt",
	} {
		src, _ := securejoin.SecureJoin(path, item)
		dst, _ := securejoin.SecureJoin(documentationPath, item)
		_ = os.Rename(src, dst)
	}

	return nil
}
