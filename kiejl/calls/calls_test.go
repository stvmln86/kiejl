package calls

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/kiejl/kiejl/items/book"
	"github.com/stvmln86/kiejl/kiejl/tools/test"
)

type MockCall struct {
	Book *book.Book
}

func NewMock(book *book.Book) (Call, error) {
	return &MockCall{book}, nil
}

func (c *MockCall) Name() string {
	return "mock"
}

func (c *MockCall) Help() string {
	return "A mock Call for unit testing."
}

func (c *MockCall) Paras() []string {
	return []string{"PARAMETER:default"}
}

func (c *MockCall) Run(w io.Writer, pairs map[string]string) error {
	fmt.Fprintf(w, "PARAMETER=%s\n", pairs["PARAMETER"])
	return nil
}

func assertCall(t *testing.T, cfun NewCallFunc) {
	call, err := cfun(nil)
	assert.NotNil(t, call)
	assert.NotEmpty(t, call.Name())
	assert.NotEmpty(t, call.Help())
	assert.NotNil(t, call.Name())
	assert.NoError(t, err)
}

func runCall(t *testing.T, argus ...string) (string, error) {
	w := bytes.NewBuffer(nil)
	dire := test.TempDire(t)
	book := book.New(dire, ".extn", 0666)
	err := Run(w, book, argus)
	return w.String(), err
}

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Calls["mock"] = NewMock

	// success
	err := Run(w, nil, []string{"mock", "argument"})
	assert.Equal(t, "PARAMETER=argument\n", w.String())
	assert.NoError(t, err)

	// error - does not exist
	err = Run(nil, nil, []string{"nope"})
	test.AssertErr(t, err, `cannot run "nope" - does not exist`)
}
