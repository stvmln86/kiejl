package book

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/kiejl/kiejl/items/note"
	"github.com/stvmln86/kiejl/kiejl/tools/test"
)

func tempBook(t *testing.T) *Book {
	dire := test.TempDire(t)
	return New(dire, ".extn", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := tempBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := tempBook(t)

	// success
	note, err := book.Create("create")
	assert.Contains(t, note.Orig, "create.extn")
	test.AssertFile(t, note.Orig, "")
	assert.NoError(t, err)
}

func TestCreateOrGet(t *testing.T) {
	// setup
	book := tempBook(t)

	// success - created
	note, err := book.CreateOrGet("create")
	assert.Contains(t, note.Orig, "create.extn")
	test.AssertFile(t, note.Orig, "")
	assert.NoError(t, err)

	// success - exists
	note, err = book.CreateOrGet("create")
	assert.Contains(t, note.Orig, "create.extn")
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := tempBook(t)

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})
	assert.Len(t, notes, 1)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := tempBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.NoError(t, err)

	// error - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	test.AssertErr(t, err, "cannot find file %q - does not exist")
}

func TestList(t *testing.T) {
	// setup
	book := tempBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Contains(t, notes[0].Orig, "alpha.extn")
	assert.Contains(t, notes[1].Orig, "bravo.extn")
}
