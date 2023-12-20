package structs

import (
	"bytes"
	"fmt"
	"hello/testhelper"
	"io"
	"testing"
)

var countDownStart = 3
var finalMessage = "Go"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, "Go")
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := &SpySleeper{}
	Countdown(&buffer, sleeper)
	if sleeper.Calls != 3 {
		t.Errorf("sleep was called %d times but expected %d", sleeper.Calls, 3)
	}
	got := buffer.String()
	wanted := `3
2
1
Go
`
	testhelper.AssertCorrectMessage(t, got, wanted)

}
