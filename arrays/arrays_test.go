package arrays

import (
	"fmt"
	"hello/testhelper"
	"slices"
	"testing"
)

func Sum(numbers []int) int {
	var sum = 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(lists ...[]int) []int {
	var rc []int
	for _, l := range lists {
		rc = append(rc, Sum(l))
	}
	return rc
}

func SumAllTails(lists ...[]int) []int {
	var rc []int
	for _, l := range lists {
		if len(l) == 0 {
			rc = append(rc, 0)
		} else {
			rc = append(rc, Sum(l[1:]))
		}
	}
	return rc
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2, 3}, []int{6, 7, 8})
	fmt.Println(got)
	// Output: [6 21]
}

func TestSlices(t *testing.T) {
	a1 := make([]int, 5)
	a1[0] = 22
	b1 := slices.Index(a1, 22)
	testhelper.AssertInteger(t, 0, b1)
	testhelper.AssertInteger(t, -1, slices.Index(a1, 33))
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum tails of 2 arrays", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{6, 7, 8})
		wanted := []int{5, 15}
		testhelper.AssertArrayIntEqual(t, got, wanted)
	})

	t.Run("Sum tails of an empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		wanted := []int{0, 0}
		testhelper.AssertArrayIntEqual(t, got, wanted)

	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{6, 7, 8})
	wanted := []int{6, 21}
	testhelper.AssertArrayIntEqual(t, got, wanted)

}

func TestSum(t *testing.T) {
	t.Run("Sum a fixed size array", func(t *testing.T) {
		numbers := []int{5, 2, 3, 4, 1}

		got := Sum(numbers)
		testhelper.AssertInteger(t, 15, got)
	})

	t.Run("Sum a variable size array", func(t *testing.T) {
		numbers := []int{5, 2, 3}

		got := Sum(numbers)
		testhelper.AssertInteger(t, 10, got)
	})

}
