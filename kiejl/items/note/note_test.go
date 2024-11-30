package note

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/kiejl/kiejl/tools/test"
)

func tempNote(t *testing.T) *Note {
	orig := test.TempFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := tempNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := tempNote(t)
	dest := strings.Replace(note.Orig, ".extn", ".trash", -1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := tempNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := tempNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := tempNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, test.MockFiles["alpha.extn"], body)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := tempNote(t)

	// success
	err := note.Update("Updated.\n")
	test.AssertFile(t, note.Orig, "Updated.\n")
	assert.NoError(t, err)
}
