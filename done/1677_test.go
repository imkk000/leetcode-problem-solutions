package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1677"
  title: matrix-diagonal-sum
  lang: golang
  type: large
  inputString: |-
    [[5]]
    [[1,1],[1,1]]
    [[1,2,3],[4,5,6],[7,8,9]]
    [[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]
*/

func diagonalSum(mat [][]int) (s int) {
	n := len(mat) - 1
	for i := 0; i <= n; i++ {
		s += mat[i][i]
		if i != n-i {
			s += mat[i][n-i]
		}
	}
	return
}

func Test_1677(t *testing.T) {
	NewTestcases(t).
		Add(5, Make2DMatrixInt("[[5]]")).
		Add(4, Make2DMatrixInt("[[1,1],[1,1]]")).
		Add(25, Make2DMatrixInt("[[1,2,3],[4,5,6],[7,8,9]]")).
		Add(8, Make2DMatrixInt("[[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]")).
		Add(55, Make2DMatrixInt("[[7,3,1,9],[3,4,6,9],[6,9,6,6],[9,5,8,5]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := diagonalSum(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
