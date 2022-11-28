package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "200"
  title: number-of-islands
  lang: golang
  type: large
  inputString: |-
    [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
    [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]
*/

func numIslands(grid [][]byte) (s int) {
	m, n := len(grid), len(grid[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				markIsland(grid, m, n, r, c)
				s++
			}
		}
	}
	return
}

func markIsland(g [][]byte, m, n, r, c int) {
	// mark visited point
	g[r][c] = '0'

	// travel each vertical point
	if r > 0 && g[r-1][c] == '1' {
		markIsland(g, m, n, r-1, c)
	}
	if r < m-1 && g[r+1][c] == '1' {
		markIsland(g, m, n, r+1, c)
	}

	// travel each horizontal point
	if c > 0 && g[r][c-1] == '1' {
		markIsland(g, m, n, r, c-1)
	}
	if c < n-1 && g[r][c+1] == '1' {
		markIsland(g, m, n, r, c+1)
	}
}

func Test_200(t *testing.T) {
	NewTestcases(t).
		Add(1, [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}).
		Add(3, [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}).
		Each(func(a *assert.Assertions, td TestData) {
			actual := numIslands(td.Input.([][]byte))

			a.Equal(td.Expectation, actual)
		})
}
