package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/kiejl/kiejl/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "create.extn")

	// success
	err := Create(dest, "Created.\n", 0666)
	test.AssertFile(t, dest, "Created.\n")
	assert.NoError(t, err)

	// error - already exists
	err = Create(dest, "Created.\n", 0666)
	assert.ErrorContains(t, err, "already exists")
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")
	dest := strings.Replace(orig, ".extn", ".trash", -1)

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// error - does not exist
	err = Delete(orig)
	assert.ErrorContains(t, err, "does not exist")
}

func TestExists(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope.extn")
	assert.False(t, ok)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, test.MockFiles["alpha.extn"], body)
	assert.NoError(t, err)

	// error - does not exist
	body, err = Read("/nope.extn")
	assert.Empty(t, body)
	assert.ErrorContains(t, err, "does not exist")
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	err := Update(orig, "Updated.\n", 0666)
	test.AssertFile(t, orig, "Updated.\n")
	assert.NoError(t, err)

	// error - does not exist
	err = Update("/nope.extn", "Updated.\n", 0666)
	assert.ErrorContains(t, err, "does not exist")
}
