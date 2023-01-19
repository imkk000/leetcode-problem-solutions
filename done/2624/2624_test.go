package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2624"
  title: difference-between-element-sum-and-digit-sum-of-an-array
  lang: golang
  type: large
  inputString: |-
    [1,15,6,3]
    [1,2,3,4]
    [1,1,2,3,4]
    [1]
    [1,11]
*/

func differenceOfSum(nums []int) int {
	var digits [2001]int
	var se, sm int
	for i, n := range nums {
		se += n
		if digits[n] == 0 {
			for nums[i] > 0 {
				sm += nums[i] % 10
				nums[i] /= 10
			}
		}
		sm += digits[n]
	}
	return abs(se - sm)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Test_2624(t *testing.T) {
	NewTestcases(t).
		Add(9, MakeIntSlice("[1,15,6,3]")).
		Add(0, MakeIntSlice("[1,2,3,4]")).
		Add(0, MakeIntSlice("[1,1,2,3,4]")).
		Add(0, MakeIntSlice("[1]")).
		Add(9, MakeIntSlice("[1,11]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := differenceOfSum(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
