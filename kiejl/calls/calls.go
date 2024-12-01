// Package calls implements the Call type and collections.
package calls

import (
	"fmt"
	"io"

	"github.com/stvmln86/kiejl/kiejl/items/book"
	"github.com/stvmln86/kiejl/kiejl/tools/clui"
)

// Call is a user-facing callable command-line function.
type Call interface {
	// Name returns the Call's callable name.
	Name() string

	// Help returns the Call's help string.
	Help() string

	// Paras returns the Call's argument parameters.
	Paras() []string

	// Run executes the Call's logic with parsed arguments.
	Run(io.Writer, map[string]string) error
}

// NewCallFunc is a function that returns a new Call.
type NewCallFunc func(*book.Book) (Call, error)

// Calls is a map of all defined Call implementations.
var Calls = map[string]NewCallFunc{
	"list": NewList,
}

// Run parses an argument slice and executes the matching Call with parsed arguments.
func Run(w io.Writer, book *book.Book, argus []string) error {
	name, argus := clui.Split(argus)
	cfun, ok := Calls[name]
	if !ok {
		return fmt.Errorf("cannot run %q - does not exist", name)
	}

	call, err := cfun(book)
	if err != nil {
		return err
	}

	pairs, err := clui.Parse(call.Paras(), argus)
	if err != nil {
		return err
	}

	return call.Run(w, pairs)
}
