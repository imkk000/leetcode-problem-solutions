package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "363"
  title: max-sum-of-rectangle-no-larger-than-k
  lang: golang
  type: large
  inputString: |-
    [[1,0,1],[0,-2,3]]
    2
    [[2,2,-1]]
    3
*/

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := range matrix {
		sum[i+1] = make([]int, n+1)
		for j := range matrix[i] {
			sum[i+1][j+1] = matrix[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}
	res, c := math.MinInt32, 0
	for r1 := 1; r1 <= m; r1++ {
		for c1 := 1; c1 <= n; c1++ {
			for r2 := r1; r2 <= m; r2++ {
				for c2 := c1; c2 <= n; c2++ {
					c = sum[r2][c2] - sum[r1-1][c2] - sum[r2][c1-1] + sum[r1-1][c1-1]
					if c <= k {
						res = max(res, c)
					}
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Test_363(t *testing.T) {
	type Data struct {
		mat [][]int
		k   int
	}
	NewTestcases(t).
		Add(2, Data{
			mat: [][]int{{1, 0, 1}, {0, -2, 3}},
			k:   2,
		}).
		Add(3, Data{
			mat: [][]int{{2, 2, -1}},
			k:   3,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := maxSumSubmatrix(input.mat, input.k)

			a.Equal(td.Expectation, actual)
		})
}
