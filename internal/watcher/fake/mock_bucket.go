package fake

import (
	"fmt"
	"net/url"
	"testing"

	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
	"gotest.tools/assert"
	"gotest.tools/fs"
)

func OpenTestFileBucket(t *testing.T) *blob.Bucket {
	t.Helper()

	dir := fs.NewDir(t, "")
	url, err := url.Parse(fmt.Sprintf("file://%s", dir.Path()))
	assert.NilError(t, err)

	secretKey := []byte("1234")
	b, err := fileblob.OpenBucket(dir.Path(), &fileblob.Options{URLSigner: fileblob.NewURLSignerHMAC(url, secretKey)})
	assert.NilError(t, err)

	t.Cleanup(func() { _ = b.Close() })

	return b
}
