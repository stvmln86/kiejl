// Package test implements unit testing mocks and functions.
package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ErrorForms is a map of error message placeholders to regular expressions.
var ErrorForms = map[string]string{
	"%q": `".*"`,
	"%s": `.*`,
	"%w": `.*`,
}

// MockFiles is a map of mock note files.
var MockFiles = map[string]string{
	"alpha.extn":    "Alpha note.\n",
	"bravo.extn":    "Bravo note.\n",
	"charlie.trash": "Charlie note (deleted).\n",
}

// AssertDire asserts a directory's contents are equal to a base:body map.
func AssertDire(t *testing.T, dire string, pairs map[string]string) {
	for base, body := range pairs {
		orig := filepath.Join(dire, base)
		bytes, err := os.ReadFile(orig)
		assert.Equal(t, body, string(bytes))
		assert.NoError(t, err)
	}
}

// AssertErr asserts an error's message matches a string template.
func AssertErr(t *testing.T, err error, text string) {
	for hold, regx := range ErrorForms {
		if strings.Contains(text, hold) {
			text = strings.Replace(text, hold, regx, -1)
		}
	}

	assert.Regexp(t, text, err.Error())
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// TempDire returns a temporary directory populated with MockFiles entries.
func TempDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockFiles {
		dest := filepath.Join(dire, base)
		os.WriteFile(dest, []byte(body), 0666)
	}

	return dire
}

// TempFile returns a MockFiles entry as a temporary file.
func TempFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	os.WriteFile(dest, []byte(MockFiles[base]), 0666)
	return dest
}
