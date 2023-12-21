package dp

import (
	"os"
	"strings"
)

func readFile(fileName string) (string, error) {
	s, err := os.ReadFile(fileName)
	var rc string
	if err == nil {
		rc = string(s)
	}
	return rc, err
}

const block = -2

func parseFile(contents string) ([][]int, [2]int) {
	rows := strings.Split(contents, "\n")
	matrix := make([][]int, len(rows))
	var start [2]int
	for i, s := range rows {
		matrix[i] = make([]int, len(s))

		for j, c := range s {
			if c == '.' {
				matrix[i][j] = -1
			} else if c == '#' {
				matrix[i][j] = block
			} else if c == 'S' {
				start[0] = i
				start[1] = j
				matrix[i][j] = 0
			}
		}
	}
	return matrix, start
}

// getNeighbours gets a list of [row, col] tuples of neighbours
func getNeighbours(mat [][]int, m, n int) [][2]int {
	width := len(mat[0])
	height := len(mat)
	if m == 0 {
		if n == 0 {
			// top left corner
			rc := [][2]int{{1, 0}, {0, 1}}
			return rc
		} else if n == width-1 {
			// top right corner
			rc := [][2]int{{1, n}, {0, n - 1}}
			return rc
		} else {
			// top row
			rc := [][2]int{{0, n - 1}, {0, n + 1}, {1, n}}
			return rc
		}
	} else if n == 0 {
		// bottom left corner
		if m == height-1 {
			rc := [][2]int{{m - 1, 0}, {m, 1}}
			return rc
		} else {
			// left side
			rc := [][2]int{{m - 1, n}, {m + 1, n}, {m, n + 1}}
			return rc
		}
	} else if n == width-1 {
		if m == height-1 {
			// bottom right corner
			rc := [][2]int{{m - 1, n}, {m, n - 1}}
			return rc
		} else {
			// right hand column
			rc := [][2]int{{m - 1, n}, {m + 1, n}, {m, n - 1}}
			return rc
		}
	} else if m == height-1 {
		// bottom row
		rc := [][2]int{{m - 1, n}, {m, n - 1}, {m, n + 1}}
		return rc
	} else {
		// we are in an interior cell
		return [][2]int{{m - 1, n}, {m, n - 1}, {m, n + 1}, {m + 1, n}}
	}
}

// Take steps one at a time
// loop over matrix, updating distance if can be reached from previous
// step
// could perhaps be faster solution iterating out from initial point.
func explore(mat_ptr *[][]int, strt_m, start_n, distance int) {
	mat := *mat_ptr
	for i := 0; i < distance; i++ {
		for j := 0; j < len(mat); j++ {
			for k := 0; k < len(mat[0]); k++ {
				if mat[j][k] == i {
					neigbours := getNeighbours(mat, j, k)
					for _, n := range neigbours {
						curr_val := mat[n[0]][n[1]]

						if curr_val != block {

							// step forwards
							mat[n[0]][n[1]] = i + 1

							// current step now not reachable
							mat[j][k] = 0
						}
					}
				}
			}
		}
	}
}
