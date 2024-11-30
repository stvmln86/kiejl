package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn(".EXTN")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tTEST.NAME!\n")
	assert.Equal(t, "test-name", name)
}

func TestPath(t *testing.T) {
	// success
	path := Path("/././dire/name.extn")
	assert.Equal(t, "/dire/name.extn", path)
}
