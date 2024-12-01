package calls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRun(t *testing.T) {
	// success
	assertCall(t, NewList)
	text, err := runCall(t, "list")
	assert.Equal(t, "alpha\nbravo\n", text)
	assert.NoError(t, err)
}
