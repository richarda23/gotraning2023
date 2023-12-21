package dp

import (
	"fmt"
	"hello/testhelper"
	"testing"
)

func TestReadFile(t *testing.T) {

	contents, _ := readFile("aco14gtestdata.txt")
	fmt.Print(contents)
	testhelper.AssertInteger(t, 131, len(contents))

}

func TestParseFile(t *testing.T) {
	distance := 64
	contents, _ := readFile("aoc21data.txt")
	matrix, _ := parseFile(contents)
	testhelper.AssertInteger(t, 131, len(matrix))
	matrix = explore(matrix, 2, 2, distance)
	for _, r := range matrix {
		fmt.Println(r)
	}
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == distance {
				count++
			}
		}
	}
	fmt.Printf("count is %d\n", count)
	testhelper.AssertInteger(t, 3751, count)

}
