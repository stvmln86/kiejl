// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a body string with trimmed whitespace and a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a lowercase file extension string with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a lowercase alphanumeric name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.In(char, unicode.Letter, unicode.Number):
			chars = append(chars, char)
		case char == '-' || char == '_' || char == '.':
			chars = append(chars, '-')
		}
	}

	return strings.ToLower(string(chars))
}

// Path returns a clean file path string.
func Path(path string) string {
	return filepath.Clean(path)
}
