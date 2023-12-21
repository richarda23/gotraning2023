package dp

import (
	"hello/testhelper"
	"testing"
)

type Cell struct {
	m, n int
}

func initM(m, n, initValue int) [][]int {
	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
	}
	if initValue != 0 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; i++ {
				rows[i][j] = initValue
			}
		}
	}
	return rows
}

type pathParams struct {
	blocked []Cell
}

// pathCount counts number of ways to traverse from 0,0 to i,j
// where m = row and n = column
func pathCount(m int, n int, optionalParams pathParams) int {
	mat := initM(m, n, 0)
	// base values
	mat[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			isBlocked := false
			for _, cell := range optionalParams.blocked {
				if cell.m == i && cell.n == j {
					mat[i][j] = 0
					isBlocked = true
					break
				}
			}
			if !isBlocked {
				if i == 0 {
					mat[i][j] = mat[i][j-1]
				} else if j == 0 {
					mat[i][j] = mat[i-1][j]
				} else {
					mat[i][j] = mat[i-1][j] + mat[i][j-1]
				}
			}
		}
	}
	return mat[m-1][n-1]
}

func TestMat(t *testing.T) {
	testhelper.AssertInteger(t, 1, pathCount(1, 1, pathParams{}))
	testhelper.AssertInteger(t, 1, pathCount(1, 8, pathParams{}))
	testhelper.AssertInteger(t, 1, pathCount(8, 1, pathParams{}))
	testhelper.AssertInteger(t, 2, pathCount(2, 2, pathParams{}))
	testhelper.AssertInteger(t, 6, pathCount(3, 3, pathParams{}))
	testhelper.AssertInteger(t, 2, pathCount(3, 3, pathParams{blocked: []Cell{{1, 1}}}))
	testhelper.AssertInteger(t, 1, pathCount(3, 3, pathParams{blocked: []Cell{{1, 1}, {1, 2}}}))
	testhelper.AssertInteger(t, 0, pathCount(3, 3, pathParams{blocked: []Cell{{1, 0}, {0, 1}}}))
}

func Benchmark(t *testing.B) {
	for i := 0; i < t.N; i++ {
		pathCount(200, 200, pathParams{})
	}
}
