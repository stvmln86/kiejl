// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"os"

	"github.com/stvmln86/kiejl/kiejl/items/note"
	"github.com/stvmln86/kiejl/kiejl/tools/file"
	"github.com/stvmln86/kiejl/kiejl/tools/neat"
	"github.com/stvmln86/kiejl/kiejl/tools/path"
)

// Book is a directory of plaintext note files.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create creates and returns a new Note in the Book.
func (b *Book) Create(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)
	if err := file.Create(dest, "", b.Mode); err != nil {
		return nil, err
	}

	return note.New(dest, b.Mode), nil
}

// CreateOrGet returns a created or existing Note in the Book.
func (b *Book) CreateOrGet(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)
	if !file.Exists(orig) {
		if err := file.Create(orig, "", b.Mode); err != nil {
			return nil, err
		}
	}

	return note.New(orig, b.Mode), nil
}

// Filter returns all existing Notes in the Book passing a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note
	for _, note := range b.List() {
		ok, err := ffun(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// Get returns an existing Note in the Book.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)
	if !file.Exists(orig) {
		return nil, fmt.Errorf("cannot find file %q - does not exist", orig)
	}

	return note.New(orig, b.Mode), nil
}

// List returns all existing Notes in the Book.
func (b *Book) List() []*note.Note {
	var notes []*note.Note
	for _, orig := range path.Glob(b.Dire, b.Extn) {
		notes = append(notes, note.New(orig, b.Mode))
	}

	return notes
}
