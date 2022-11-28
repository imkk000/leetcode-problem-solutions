package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "777"
  title: toeplitz-matrix
  lang: golang
  type: large
  inputString: |-
    [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
    [[1,2],[2,2]]
*/

func isToeplitzMatrix(mat [][]int) bool {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		b := mat[i][0]
		for j := 0; j < n; j++ {
			if i+j >= m {
				break
			}
			if mat[i+j][j] != b {
				return false
			}
		}
	}
	for i := 1; i < n; i++ {
		b := mat[0][i]
		for j := 0; j < m; j++ {
			if i+j >= n {
				break
			}
			if mat[j][i+j] != b {
				return false
			}
		}
	}
	return true
}

func Test_777(t *testing.T) {
	NewTestcases(t).
		Add(true, Make2DMatrixInt("[[1,2,3,4],[5,1,2,3],[9,5,1,2]]")).
		Add(false, Make2DMatrixInt("[[1,2],[2,2]]")).
		Add(true, Make2DMatrixInt("[[1]]")).
		Add(true, Make2DMatrixInt("[[1],[2]]")).
		Add(true, Make2DMatrixInt("[[1,2]]")).
		Add(true, Make2DMatrixInt("[[1,2],[2,1]]")).
		Add(true, Make2DMatrixInt("[[1],[2],[3]]")).
		Add(true, Make2DMatrixInt("[[1,2,3]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isToeplitzMatrix(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
