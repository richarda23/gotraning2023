package iteration

import (
	"hello/testhelper"
	"strings"
	"testing"
)

func Repeat(s string, n int) string {
	var r = ""

	for i := 0; i < n; i++ {
		r += s
	}
	return r
}

func RepeatStl(s string, n int) string {
	return strings.Repeat(s, n)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 500)
	}
}

func BenchmarkRepeatStl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStl("b", 500)
	}
}

func TestIteration(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"
	testhelper.AssertCorrectMessage(t, expected, got)

}
