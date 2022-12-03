package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1"
  title: two-sum
  lang: golang
  type: small
  inputString: |-
    [2,7,11,15]
    9
    [3,2,4]
    6
    [3,3]
    6
*/

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func Test_1(t *testing.T) {
	type Data struct {
		nums   []int
		target int
	}
	NewTestcases(t).
		AddExpectation([]int{0, 1}).
		AddInput(Data{
			nums:   MakeIntSlice("[2,7,11,15]"),
			target: 9,
		}).
		// AddExpectation([]int{1, 2}).
		// AddInput(Data{
		// 	nums:   MakeIntSlice("[3,2,4]"),
		// 	target: 6,
		// }).
		// AddExpectation([]int{0, 1}).
		// AddInput(Data{
		// 	nums:   MakeIntSlice("[3,3]"),
		// 	target: 6,
		// }).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			nums, target := input.nums, input.target
			actual := twoSum(nums, target)

			a.Equal(td.Expectation, actual)
		})
}
