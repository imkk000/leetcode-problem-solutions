package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "861"
  title: flipping-an-image
  lang: golang
  type: large
  inputString: |-
    [[1]]
    [[1,1,0],[1,0,1],[0,0,0]]
    [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
*/

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	m := n
	if n&1 == 1 {
		m++
	}
	m /= 2
	for i := range image {
		for j := 0; j < m; j++ {
			image[i][j], image[i][n-j-1] = image[i][n-j-1]^1, image[i][j]^1
		}
	}
	return image
}

func Test_861(t *testing.T) {
	NewTestcases(t).
		Add(Make2DMatrixInt("[[0]]"), Make2DMatrixInt("[[1]]")).
		Add(Make2DMatrixInt("[[1,0,0],[0,1,0],[1,1,1]]"), Make2DMatrixInt("[[1,1,0],[1,0,1],[0,0,0]]")).
		Add(Make2DMatrixInt("[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]"), Make2DMatrixInt("[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := flipAndInvertImage(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
