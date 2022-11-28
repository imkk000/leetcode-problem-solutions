package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "219"
  title: contains-duplicate-ii
  lang: golang
  type: large
  inputString: |-
    [1,2,3,1]
    3
    [1,0,1,1]
    1
    [1,2,3,1,2,3]
    2
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	abs := func(n int) int {
		if n < 0 {
			return n * -1
		}
		return n
	}
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j <= i+k && j < l; j++ {
			if nums[i] == nums[j] {
				if abs(i-j) <= k {
					return true
				}
			}
		}
	}
	return false
}

func Test_219(t *testing.T) {
	type Data struct {
		nums []int
		k    int
	}

	NewTestcases(t).
		Add(true, Data{
			nums: MakeIntSlice("[1,2,3,1]"),
			k:    3,
		}).
		Add(true, Data{
			nums: MakeIntSlice("[1,0,1,1]"),
			k:    1,
		}).
		Add(false, Data{
			nums: MakeIntSlice("[1,2,3,1,2,3]"),
			k:    2,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := containsNearbyDuplicate(input.nums, input.k)

			a.Equal(td.Expectation, actual)
		})
}
