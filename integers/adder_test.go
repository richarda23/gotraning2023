package integers

import (
	"fmt"
	"hello/testhelper"
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
	testhelper.AssertInteger(t, wanted, sum)
}
