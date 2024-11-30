package test

import (
	"errors"
	"testing"
)

func TestAssertDire(t *testing.T) {
	// setup
	dire := TempDire(t)

	// success
	AssertDire(t, dire, MockFiles)
}

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New(`test "quote" - error`)

	// success
	AssertErr(t, err, "test %q - %w")
}

func TestAssertFile(t *testing.T) {
	// setup
	orig := TempFile(t, "alpha.extn")

	// success
	AssertFile(t, orig, MockFiles["alpha.extn"])
}

func TestTempDire(t *testing.T) {
	// success
	dire := TempDire(t)
	AssertDire(t, dire, MockFiles)
}

func TestTempFile(t *testing.T) {
	// success
	orig := TempFile(t, "alpha.extn")
	AssertFile(t, orig, MockFiles["alpha.extn"])
}
