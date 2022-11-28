package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "118"
  title: pascals-triangle
  lang: golang
  type: large
  inputString: |-
    1
    2
    3
    4
    5
    6
*/

func generate(r int) [][]int {
	v := make([][]int, r)
	for i := 0; i < r; i++ {
		v[i] = make([]int, i+1)
		v[i][0], v[i][i] = 1, 1
		for j := 1; j <= i/2; j++ {
			v[i][j] = v[i-1][j] + v[i-1][j-1]
			v[i][i-j] = v[i][j]
		}
	}
	return v
}

func Test_118(t *testing.T) {
	NewTestcases(t).
		Add(Make2DMatrixInt("[[1]]"), 1).
		Add(Make2DMatrixInt("[[1],[1,1]]"), 2).
		Add(Make2DMatrixInt("[[1],[1,1],[1,2,1]]"), 3).
		Add(Make2DMatrixInt("[[1],[1,1],[1,2,1],[1,3,3,1]]"), 4).
		Add(Make2DMatrixInt("[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]"), 5).
		Add(Make2DMatrixInt("[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1],[1,5,10,10,5,1]]"), 6).
		Each(func(a *assert.Assertions, td TestData) {
			actual := generate(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
