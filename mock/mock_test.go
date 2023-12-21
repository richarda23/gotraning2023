package structs

import (
	"bytes"
	"fmt"
	"hello/testhelper"
	"io"
	"reflect"
	"testing"
	"time"
)

var countDownStart = 3
var finalMessage = "Go"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, "Go")
}

func TestConfigurableSleeper(t *testing.T) {
	stime := 5 * time.Second
	spytime := &SpyTime{}
	sleeper := ConfigurableSleeper{stime, spytime.Sleep}
	sleeper.Sleep()
	testhelper.AssertInteger(t, 5, int(spytime.durationSlept.Seconds()))
}

func TestGreet(t *testing.T) {

	t.Run("prints correct answers", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := &SpySleeper{}
		Countdown(&buffer, sleeper)
		if len(sleeper.Calls) != 3 {
			t.Errorf("sleep was called %d times but expected %d", len(sleeper.Calls), 3)
		}
		got := buffer.String()
		wanted := `3
2
1
Go
`
		testhelper.AssertCorrectMessage(t, got, wanted)
	})

	t.Run("prints correct sequence", func(t *testing.T) {
		expected := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		sleeper := &SpySleeper{}
		Countdown(sleeper, sleeper)
		if !reflect.DeepEqual(expected, sleeper.Calls) {
			t.Errorf("wanted %v but got %v", expected, sleeper.Calls)
		}
	})

}
