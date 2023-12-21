package structs

import (
	"bytes"
	"fmt"
	"hello/testhelper"
	"io"
	"testing"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s", name)
}

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "Bob")
	got := b.String()
	testhelper.AssertCorrectMessage(t, "Hello Bob", got)
}
