package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "48"
  title: rotate-image
  lang: golang
  type: large
  inputString: |-
    [[2]]
    [[1,2],[4,3]]
    [[1,2,3],[4,5,6],[7,8,9]]
    [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
    [[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]
    [[2,29,20,26,16,28],[12,27,9,25,13,21],[32,33,32,2,28,14],[13,14,32,27,22,26],[33,1,20,7,21,7],[4,24,1,6,32,34]]
*/

func rotate(m [][]int) {
	l := len(m)
	for c, n := 0, l-1; c < l/2; c, n = c+1, n-1 {
		for i := 0; i < n-c; i++ {
			m[c][c+i], m[c+i][n], m[n][n-i], m[n-i][c] =
				m[n-i][c], m[c][c+i], m[c+i][n], m[n][n-i]
		}
	}
}

func Test_48(t *testing.T) {
	NewTestcases(t).
		AddExpectation(Make2DMatrixInt("[[2]]")).
		AddInput(Make2DMatrixInt("[[2]]")).
		AddExpectation(Make2DMatrixInt("[[4,1],[3,2]]")).
		AddInput(Make2DMatrixInt("[[1,2],[4,3]]")).
		AddExpectation(Make2DMatrixInt("[[7,4,1],[8,5,2],[9,6,3]]")).
		AddInput(Make2DMatrixInt("[[1,2,3],[4,5,6],[7,8,9]]")).
		AddExpectation(Make2DMatrixInt("[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]")).
		AddInput(Make2DMatrixInt("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]")).
		AddExpectation(Make2DMatrixInt("[[21,16,11,6,1],[22,17,12,7,2],[23,18,13,8,3],[24,19,14,9,4],[25,20,15,10,5]]")).
		AddInput(Make2DMatrixInt("[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]")).
		AddExpectation(Make2DMatrixInt("[[4,33,13,32,12,2],[24,1,14,33,27,29],[1,20,32,32,9,20],[6,7,27,2,25,26],[32,21,22,28,13,16],[34,7,26,14,21,28]]")).
		AddInput(Make2DMatrixInt("[[2,29,20,26,16,28],[12,27,9,25,13,21],[32,33,32,2,28,14],[13,14,32,27,22,26],[33,1,20,7,21,7],[4,24,1,6,32,34]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := td.Input.([][]int)
			rotate(actual)

			a.Equal(td.Expectation, actual)
		})
}
