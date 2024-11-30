// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/stvmln86/kiejl/kiejl/tools/file"
	"github.com/stvmln86/kiejl/kiejl/tools/neat"
	"github.com/stvmln86/kiejl/kiejl/tools/path"
)

// Note is a plaintext note file in a directory.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete renames the Note's extension to ".trash".
func (n *Note) Delete() error {
	return file.Delete(n.Orig)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Name returns the Note's base name without the extension.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return neat.Body(body), err
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	return file.Update(n.Orig, body, n.Mode)
}
