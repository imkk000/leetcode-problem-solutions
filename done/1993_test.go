package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1993"
  title: sum-of-all-subset-xor-totals
  lang: golang
  type: large
  inputString: |-
    [3]
    [1,3]
    [5,1,6]
    [1,2,3,4]
    [1,2,3,4,5]
    [1,2,3,4,5,6]
    [1,2,3,4,5,6,7]
    [1,2,3,4,5,6,7,8]
    [1,2,3,4,5,6,7,8,9]
    [1,2,3,4,5,6,7,8,9,10]
    [1,2,3,4,5,6,7,8,9,10,11]
    [1,2,3,4,5,6,7,8,9,10,11,12]
*/

func subsetXORSum(nums []int) int {
	return dfs(nums, len(nums), 0, 0)
}

func dfs(nums []int, n, s, x int) int {
	c := x
	for i := s; i < n; i++ {
		c += dfs(nums, n, i+1, x^nums[i])
	}
	return c
}

func Test_1993(t *testing.T) {
	NewTestcases(t).
		Add(6, MakeIntSlice("[1,3]")).
		Add(28, MakeIntSlice("[5,1,6]")).
		Add(480, MakeIntSlice("[3,4,5,6,7,8]")).
		Add(56, MakeIntSlice("[1,2,3,4]")).
		Add(112, MakeIntSlice("[1,2,3,4,5]")).
		Add(224, MakeIntSlice("[1,2,3,4,5,6]")).
		Add(448, MakeIntSlice("[1,2,3,4,5,6,7]")).
		Add(1920, MakeIntSlice("[1,2,3,4,5,6,7,8]")).
		Add(3840, MakeIntSlice("[1,2,3,4,5,6,7,8,9]")).
		Add(7680, MakeIntSlice("[1,2,3,4,5,6,7,8,9,10]")).
		Add(15360, MakeIntSlice("[1,2,3,4,5,6,7,8,9,10,11]")).
		Add(30720, MakeIntSlice("[1,2,3,4,5,6,7,8,9,10,11,12]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := subsetXORSum(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
