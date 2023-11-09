package unpack

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-logr/logr"
	"github.com/mholt/archiver/v3"
	"go.artefactual.dev/tools/temporal"
)

// FormatList represents a list of file formats.
type formatList struct {
	items []fileFormat
}

// FileFormat represents a file format.
type fileFormat struct {
	PUID string
	Name string
}

// An Unarchiver extracts the src archive's contents to dest.
type Unarchiver interface {
	Unarchive(src string, dest string) error
}

// ArchiverCmd represents the mholt/archiver unarchive command.
type ArchiverCmd struct{}

// UnpackActivity represents an Unpack Activity instantiation.
type UnpackActivity struct {
	UnarchiveImpl Unarchiver
	logger        logr.Logger
}

// UnpackParams represents the arguments passed to UnpackActivity.
type UnpackParams struct {
	PackagePath string
}

// UnpackResults represents the results returned from UnpackActivity.
type UnpackResults struct {
	PackagePath string
	ExtractPath string
}

// AllowedFileFormats lists the https://www.nationalarchives.gov.uk/PRONOM
// PUIDs of the file archive types accepted by sdps_preprocessor.
var allowedFormats = formatList{
	items: []fileFormat{
		{"x-fmt/263", "zip"},
		{"x-fmt/265", "tar"},
	},
}

// Unpack extracts the archive package at srcPath to ExtractPath, and returns
// the PackagePath and ExtractPath.
func (u *UnpackActivity) Unpack(ctx context.Context, args UnpackParams) (UnpackResults, error) {
	u.logger = temporal.GetLogger(ctx)
	u.logger.V(1).Info("Unpack Activity called.", "path", args.PackagePath)

	res := UnpackResults{
		PackagePath: args.PackagePath,
	}

	iface, err := archiver.ByExtension(args.PackagePath)
	if err != nil {
		return res, temporal.NewNonRetryableError(fmt.Errorf("couldn't find a decompressor for %s: %v", args.PackagePath, err))
	}
	unar, ok := iface.(archiver.Unarchiver)
	if !ok {
		return res, temporal.NewNonRetryableError(fmt.Errorf("couldn't find a decompressor for %s", args.PackagePath))
	}

	if err := unar.Unarchive(args.PackagePath, filepath.Dir(args.PackagePath)); err != nil {
		return res, temporal.NewNonRetryableError(fmt.Errorf("error extracting file %s: %v", args.PackagePath, err))
	}

	res.ExtractPath = filepath.Dir(args.PackagePath)
	return res, nil
}

// TrimExtension returns path with the final file extension removed.
func trimExtension(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

// NewArchiverCmd returns a new archiver command wrapper.
func NewArchiverCmd() ArchiverCmd {
	return ArchiverCmd{}
}

// Unarchive extracts the src archive contents to dest using the archiver package.
func (c ArchiverCmd) Unarchive(src string, dest string) error {
	if err := archiver.Unarchive(src, dest); err != nil {
		return err
	}

	return nil
}
