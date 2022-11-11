package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "26"
  title: remove-duplicates-from-sorted-array
  lang: golang
  type: large
  inputString: |-
    [1,1,2]
    [0,0,1,1,1,2,2,3,3,4]
*/

func removeDuplicates(nums []int) int {
	c := 1
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[c-1] != nums[i] {
			nums[c] = nums[i]
			c++
		}
	}
	return c
}

func Test_26(t *testing.T) {
	NewTestcases(t).
		Add(MakeIntSlice("[1,2]"), MakeIntSlice("[1,1,2]")).
		Add(MakeIntSlice("[0,1,2,3,4]"), MakeIntSlice("[0,0,1,1,1,2,2,3,3,4]")).
		Add(MakeIntSlice("[1,2,3,4]"), MakeIntSlice("[1,1,1,2,2,3,4,4,4,4]")).
		Add(MakeIntSlice("[1]"), MakeIntSlice("[1]")).
		Add(MakeIntSlice("[-100, -99, 0]"), MakeIntSlice("[-100,-100,-99,0,0,0,0,0]")).
		Each(func(a *assert.Assertions, td TestData) {
			expectation := td.Expectation.([]int)
			input := td.Input.([]int)
			actual := removeDuplicates(input)

			a.Equal(len(expectation), actual)
			a.Equal(td.Expectation, input[:len(expectation)])
		})
}
