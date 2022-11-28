package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1916"
  title: find-center-of-star-graph
  lang: golang
  type: large
  inputString: |-
    [[1,2],[2,3],[4,2]]
    [[1,2],[5,1],[1,3],[1,4]]
*/

func findCenter(edges [][]int) int {
	if edges[0][0]^edges[1][0] == 0 || edges[0][0]^edges[1][1] == 0 {
		return edges[0][0]
	}
	return edges[0][1]
}

func Test_1916(t *testing.T) {
	NewTestcases(t).
		Add(2, Make2DMatrixInt("[[1,2],[2,3],[4,2]]")).
		Add(1, Make2DMatrixInt("[[1,2],[5,1],[1,3],[1,4]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := findCenter(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
