package leetcode_test

import (
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1253"
  title: sort-the-matrix-diagonally
  lang: golang
  type: large
  inputString: |-
    [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
    [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]
    [[1,2],[3,4]]
    [[1,2,3,4],[4,5,6,7]]
    [[1,2],[3,4],[4,5],[7,8]]
    [[5,4,3],[3,4,2],[7,9,1]]
    [[3,2,1]]
    [[2],[1],[4]]
    [[1]]
*/

func max1253(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min1253(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	so := max1253(m, n) - 1
	var v1, v2 []int
	for i := 0; i < so; i++ {
		lm, ln := min1253(n-i, m), min1253(m-i, n)
		si := max1253(lm, ln)
		if lm > 0 {
			v1 = make([]int, lm)
		}
		if ln > 0 {
			v2 = make([]int, ln)
		}
		for j := 0; j < si; j++ {
			if j < lm {
				v1[j] = mat[j][i+j]
			}
			if j < ln {
				v2[j] = mat[i+j][j]
			}
		}
		sort.Ints(v1)
		sort.Ints(v2)
		for j := 0; j < si; j++ {
			if j < lm {
				mat[j][i+j] = v1[j]
			}
			if j < ln {
				mat[i+j][j] = v2[j]
			}
		}
	}
	return mat
}

func Test_1253(t *testing.T) {
	NewTestcases(t).
		AddExpectation(NewMatrixFromStr("[[1,1,1,1],[1,2,2,2],[1,2,3,3]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[3,3,1,1],[2,2,1,2],[1,1,1,2]]")).
		AddExpectation(NewMatrixFromStr("[[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66],[84,28,75,33,55,68]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]")).
		AddExpectation(NewMatrixFromStr("[[1,2],[3,4]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[1,2],[3,4]]")).
		AddExpectation(NewMatrixFromStr("[[1,2,3,4],[4,5,6,7]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[1,2,3,4],[4,5,6,7]]")).
		AddExpectation(NewMatrixFromStr("[[1,2],[3,4],[4,5],[7,8]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[1,2],[3,4],[4,5],[7,8]]")).
		AddExpectation(NewMatrixFromStr("[[1,2,3],[3,4,4],[7,9,5]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[5,4,3],[3,4,2],[7,9,1]]")).
		AddExpectation(NewMatrixFromStr("[[3,2,1]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[3,2,1]]")).
		AddExpectation(NewMatrixFromStr("[[2],[1],[4]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[2],[1],[4]]")).
		AddExpectation(NewMatrixFromStr("[[1]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[1]]")).
		AddExpectation(NewMatrixFromStr("[[37,98,82,45,42]]").GetMatrix()).
		AddInput(NewMatrixFromStr("[[37,98,82,45,42]]")).
		Each(func(a *assert.Assertions, td TestData) {
			mat := td.Input.(*Matrix)
			actual := diagonalSort(mat.GetMatrix())

			a.Equal(td.Expectation, actual)
		})
}
