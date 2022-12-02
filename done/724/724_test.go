package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "724"
  title: find-pivot-index
  lang: golang
  type: large
  inputString: |-
    [1,7,3,6,5,6]
    [1,2,3]
    [2,1,-1]
    [1]
*/

func pivotIndex(v []int) int {
	var sum int
	for i := range v {
		sum += v[i]
	}
	left, right := 0, sum
	for i := range v {
		right -= v[i]
		if left == right {
			return i
		}
		left += v[i]
	}
	return -1
}

func Test_724(t *testing.T) {
	NewTestcases(t).
		Add(3, MakeIntSlice("[1,7,3,6,5,6]")).
		Add(-1, MakeIntSlice("[1,2,3]")).
		Add(0, MakeIntSlice("[2,1,-1]")).
		Add(0, MakeIntSlice("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := pivotIndex(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
