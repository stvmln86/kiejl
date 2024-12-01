// Package main implements the main program functions.
package main

import (
	"fmt"
	"os"

	"github.com/stvmln86/kiejl/kiejl/calls"
	"github.com/stvmln86/kiejl/kiejl/items/book"
	"github.com/stvmln86/kiejl/kiejl/tools/clui"
)

// try prints and exits on a non-nil error.
func try(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

// main runs the main Kiejl program.
func main() {
	dire, err := clui.GetEnv("KIEJL_DIR")
	try(err)

	extn, err := clui.GetEnv("KIEJL_EXT")
	try(err)

	book := book.New(dire, extn, 0666)
	try(calls.Run(os.Stdout, book, os.Args[1:]))
}
