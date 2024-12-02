package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/kiejl/kiejl/tools/test"
)

func TestGetEnv(t *testing.T) {
	// setup
	os.Setenv("NAME", "Value.\n")
	os.Setenv("BLANK", "\n")

	// success
	evar, err := GetEnv("NAME")
	assert.Equal(t, "Value.", evar)
	assert.NoError(t, err)

	// error - is not set
	evar, err = GetEnv("NOPE")
	assert.Empty(t, evar)
	test.AssertErr(t, err, `environment variable "NOPE" is not set`)

	// error - is blank
	evar, err = GetEnv("BLANK")
	assert.Empty(t, evar)
	test.AssertErr(t, err, `environment variable "BLANK" is blank`)
}

func TestParse(t *testing.T) {
	// success - real value
	pairs, err := Parse([]string{"PARAMETER"}, []string{"argument"})
	assert.Equal(t, map[string]string{"PARAMETER": "argument"}, pairs)
	assert.NoError(t, err)

	// success - default value
	pairs, err = Parse([]string{"PARAMETER:default"}, nil)
	assert.Equal(t, map[string]string{"PARAMETER": "default"}, pairs)
	assert.NoError(t, err)

	// success - default empty value
	pairs, err = Parse([]string{"PARAMETER:"}, nil)
	assert.Equal(t, map[string]string{"PARAMETER": ""}, pairs)
	assert.NoError(t, err)

	// error - missing argument
	pairs, err = Parse([]string{"PARAMETER"}, nil)
	assert.Nil(t, pairs)
	test.AssertErr(t, err, `missing "PARAMETER" argument`)
}

func TestSplit(t *testing.T) {
	// success - zero arguments
	name, argus := Split(nil)
	assert.Empty(t, name)
	assert.Nil(t, argus)

	// success - one argument
	name, argus = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Nil(t, argus)

	// success - multiple arguments
	name, argus = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, argus)
}
