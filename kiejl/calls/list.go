package calls

import (
	"fmt"
	"io"

	"github.com/stvmln86/kiejl/kiejl/items/book"
)

// List is a Call that lists existing Notes.
type List struct {
	Book *book.Book
}

// NewList returns a new List.
func NewList(book *book.Book) (Call, error) {
	return &List{book}, nil
}

// Name returns the List's callable name.
func (c *List) Name() string {
	return "list"
}

// Help returns the List's help string.
func (c *List) Help() string {
	return "List all existing Notes."
}

// Paras returns the List's argument parameters.
func (c *List) Paras() []string {
	return []string{}
}

// Run executes the List's logic with parsed arguments.
func (c *List) Run(w io.Writer, pairs map[string]string) error {
	for _, note := range c.Book.List() {
		fmt.Fprintf(w, "%s\n", note.Name())
	}

	return nil
}
