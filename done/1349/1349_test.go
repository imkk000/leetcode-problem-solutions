package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1349"
  title: check-if-it-is-a-straight-line
  lang: golang
  type: large
  inputString: |-
    [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
    [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
    [[1,1],[0,1]]
    [[-1,0],[0,0]]
    [[0,0],[1,1],[9,9]]
    [[0,0],[0,1],[0,-1]]
    [[0,0],[0,5],[5,5],[5,0]]
*/

func checkStraightLine(c [][]int) bool {
	m := calculateSlope(c[0][0], c[0][1], c[1][0], c[1][1])
	l := len(c)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if m != calculateSlope(c[i][0], c[i][1], c[j][0], c[j][1]) {
				return false
			}
		}
	}
	return true
}

func calculateSlope(x1, y1, x2, y2 int) int {
	if x := x2 - x1; x != 0 {
		return (y2 - y1) / x
	}
	return 0
}

func Test_1349(t *testing.T) {
	NewTestcases(t).
		Add(true, Make2DMatrixInt("[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]")).
		Add(false, Make2DMatrixInt("[[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]")).
		Add(true, Make2DMatrixInt("[[1,1],[0,1]]")).
		Add(true, Make2DMatrixInt("[[-1,0],[0,0]]")).
		Add(true, Make2DMatrixInt("[[0,0],[1,1],[9,9]]")).
		Add(true, Make2DMatrixInt("[[0,0],[0,1],[0,-1]]")).
		Add(false, Make2DMatrixInt("[[0,0],[0,5],[5,5],[5,0]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := checkStraightLine(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
