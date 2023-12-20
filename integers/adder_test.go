package integers

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	wanted := 4
	AssertInteger(t, wanted, sum)
}

func AssertInteger(t testing.TB, wanted, got int) {
	t.Helper()
	if wanted != got {
		t.Errorf("Wanted %d but got %d", wanted, got)
	}
}
