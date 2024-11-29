// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Base returns a path's base name with the extension.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's file extension with the leading dot.
func Extn(orig string) string {
	return filepath.Ext(orig)
}

// Glob returns all paths in a directory matching an extension.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	return origs
}

// Join returns a joined path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Name returns a path's base name without the extension.
func Name(orig string) string {
	base := filepath.Base(orig)
	extn := filepath.Ext(orig)
	return strings.TrimSuffix(base, extn)
}
