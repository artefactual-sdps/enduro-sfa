package unpack

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go.temporal.io/sdk/testsuite"
	"gotest.tools/v3/assert"
	tfs "gotest.tools/v3/fs"

	"gitlab.artefactual.com/clients/ste/sdps_preprocessing/pkg/fformat"
)

type (
	fakeUnarchiveImpl struct{}
	fakeFormatIDImpl  struct{}
)

const transferName = "ABCD202211140947---EAS-A---15.2"

// Wmod represents the wanted file modes for our test results.
var (
	wmod = map[string]fs.FileMode{
		"dir":  fs.FileMode(0o755),
		"file": fs.FileMode(0o640),
	}

	arcFiles = []struct {
		Name string
		Body string
		Mode int64
	}{
		{Name: transferName + "/", Mode: int64(wmod["dir"])},
		{
			Name: transferName + "/readme.txt",
			Body: "This archive contains some text files.",
			Mode: int64(wmod["file"]),
		},
		{
			Name: transferName + "/gopher.txt",
			Body: "Gopher names:\nGeorge\nGonzo\nGrace",
			Mode: int64(wmod["file"]),
		},
	}
)

func (f fakeUnarchiveImpl) Unarchive(src string, dest string) error {
	switch filepath.Ext(src) {
	case ".tar", ".zip":
		return nil
	}

	return fmt.Errorf("invalid archive: %q", src)
}

func testFileFormat(puid, format string) *fformat.FileFormat {
	return &fformat.FileFormat{
		ID:         puid,
		CommonName: format,
	}
}

func (f fakeFormatIDImpl) Identify(path string) (*fformat.FileFormat, error) {
	switch filepath.Ext(path) {
	case ".tar":
		return testFileFormat("x-fmt/265", "tar"), nil
	case ".txt":
		return testFileFormat("x-fmt/111", "plain text"), nil
	case ".zip":
		return testFileFormat("x-fmt/263", "zip"), nil
	}

	return &fformat.FileFormat{}, errors.New("unknown format")
}

func (f fakeFormatIDImpl) Version() string {
	return "0.1.0"
}

func createTestDir(t *testing.T) *tfs.Dir {
	return tfs.NewDir(t, "test", tfs.WithMode(wmod["dir"]))
}

func createTestTxt(t *testing.T) string {
	d := createTestDir(t)
	fp := d.Join("a.txt")
	if err := os.WriteFile(
		fp, []byte("This is a text file."), wmod["file"],
	); err != nil {
		t.Fatal(err)
	}

	return fp
}

func createTestTar(t *testing.T) string {
	d := createTestDir(t)
	fp := d.Join(transferName + ".tar")

	fh, err := os.Create(fp)
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	tw := tar.NewWriter(fh)
	for _, file := range arcFiles {
		hdr := &tar.Header{
			Name:   file.Name,
			Mode:   file.Mode,
			Size:   int64(len(file.Body)),
			Format: tar.FormatGNU,
		}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	return fp
}

func createTestZip(t *testing.T) string {
	td := createTestDir(t)
	fp := td.Join(transferName + ".zip")

	fh, err := os.Create(fp)
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	w := zip.NewWriter(fh)
	for _, file := range arcFiles {
		if file.Body != "" {
			fh := zip.FileHeader{
				Name: file.Name,
			}
			fh.SetMode(wmod["file"])

			f, err := w.CreateHeader(&fh)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := f.Write([]byte(file.Body)); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	return fp
}

func wantManifest(t *testing.T) *tfs.Manifest {
	m := tfs.Expected(t, tfs.WithMode(wmod["dir"]),
		tfs.WithDir(transferName, tfs.WithMode(wmod["dir"]),
			tfs.WithFile(
				"readme.txt",
				"This archive contains some text files.",
				tfs.WithMode(wmod["file"]),
			),
			tfs.WithFile(
				"gopher.txt",
				"Gopher names:\nGeorge\nGonzo\nGrace",
				tfs.WithMode(wmod["file"]),
			),
		),
	)

	return &m
}

func TestUnpackActivity(t *testing.T) {
	t.Parallel()

	td := createTestDir(t)
	tar := createTestTar(t)
	txt := createTestTxt(t)
	zip := createTestZip(t)

	tests := []struct {
		name    string
		args    UnpackParams
		want    UnpackResults
		wantErr string
	}{
		{
			name: "Returns path to unpacked tar contents",
			args: UnpackParams{PackagePath: tar},
			want: UnpackResults{
				PackagePath: tar,
				ExtractPath: filepath.Join(filepath.Dir(tar), transferName),
			},
			wantErr: "",
		},
		{
			name: "Returns path to unpacked zip contents",
			args: UnpackParams{PackagePath: zip},
			want: UnpackResults{
				PackagePath: zip,
				ExtractPath: filepath.Join(filepath.Dir(zip), transferName),
			},
			wantErr: "",
		},
		{
			name:    "Returns an error when srcPath is empty",
			args:    UnpackParams{PackagePath: ""},
			want:    UnpackResults{},
			wantErr: "unpack: couldn't identify file format of \"\"",
		},
		{
			name:    "Returns an error for a disallowed file format",
			args:    UnpackParams{PackagePath: txt},
			want:    UnpackResults{},
			wantErr: fmt.Sprintf("unpack: %q has an invalid file format; allowed formats are (zip, tar)", txt),
		},
		{
			name:    "Returns a file not found error",
			args:    UnpackParams{PackagePath: td.Join("missing.txt")},
			want:    UnpackResults{},
			wantErr: fmt.Sprintf("unpack: %q has an invalid file format; allowed formats are (zip, tar)", td.Join("missing.txt")),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create an empty extract dir that should be deleted before
			// extracting the archive.
			if tt.args.PackagePath != "" {
				if err := os.Mkdir(
					trimExtension(tt.args.PackagePath), wmod["dir"],
				); err != nil {
					t.Fatalf("couldn't create extract dir: %v", err)
				}
			}

			ts := &testsuite.WorkflowTestSuite{}
			env := ts.NewTestActivityEnvironment()
			activity := UnpackActivity{
				UnarchiveImpl: fakeUnarchiveImpl{},
				FormatIDImpl:  fakeFormatIDImpl{},
			}
			env.RegisterActivity(activity.Unpack)
			r, err := env.ExecuteLocalActivity(activity.Unpack, tt.args)

			if tt.wantErr == "" {
				assert.NilError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.wantErr)

				return
			}

			var got UnpackResults
			_ = r.Get(&got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackActivity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArchiverCmd(t *testing.T) {
	t.Parallel()

	t.Run("NewArchive returns an Archive", func(t *testing.T) {
		c := NewArchiverCmd()
		assert.Equal(t, c, ArchiverCmd{})
	})
}

func TestArchive_Unarchive(t *testing.T) {
	t.Parallel()

	td := createTestDir(t)
	tar := createTestTar(t)
	zip := createTestZip(t)
	noarc := func() string {
		n := "noarc.txt"
		d := tfs.NewDir(t, "noarc", tfs.WithMode(wmod["dir"]),
			tfs.WithFile(n, "I'm not an archive!", tfs.WithMode(wmod["file"])),
		)
		return d.Join(n)
	}()

	type args struct {
		dest string
		src  string
	}
	tests := []struct {
		name     string
		args     args
		wantDest string
		wantMfst *tfs.Manifest
		wantErr  string
	}{
		{
			name: "Unarchives tar to dest directory",
			args: args{
				src:  tar,
				dest: trimExtension(tar),
			},
			wantDest: td.Join(transferName),
			wantMfst: wantManifest(t),
			wantErr:  "",
		},
		{
			name: "Unarchives zip to dest directory",
			args: args{
				src:  zip,
				dest: trimExtension(zip),
			},
			wantDest: td.Join(transferName),
			wantMfst: wantManifest(t),
			wantErr:  "",
		},
		{
			name: "Return an error if archive doesn't exist",
			args: args{
				src:  td.Join("missing.tar"),
				dest: td.Join("missing"),
			},
			wantDest: "",
			wantMfst: nil,
			wantErr: fmt.Sprintf(
				"opening source archive: open %s: no such file or directory",
				td.Join("missing.tar"),
			),
		},
		{
			name: "Returns an error if archive contents can't be extracted",
			args: args{
				src:  noarc,
				dest: trimExtension(noarc),
			},
			wantDest: "",
			wantMfst: nil,
			wantErr: fmt.Sprintf(
				"format unrecognized by filename: %s", noarc,
			),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := NewArchiverCmd()
			err := c.Unarchive(tt.args.src, tt.args.dest)
			if tt.wantErr == "" {
				assert.NilError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.wantErr)
			}

			if tt.wantMfst != nil {
				assert.Assert(t, tfs.Equal(tt.args.dest, *tt.wantMfst))
			}
		})
	}
}
