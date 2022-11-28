package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "645"
  title: set-mismatch
  lang: golang
  type: large
  inputString: |-
    [1,2,2,4]
    [1,1]
    [1,1,2]
    [1,3,2,2]
*/

func findErrorNums(nums []int) []int {
	l := len(nums)
	v := make([]bool, l)
	var first, last int
	for _, n := range nums {
		if v[n-1] {
			first = n
			continue
		}
		v[n-1] = true
	}
	for i, b := range v {
		if !b {
			last = i + 1
			break
		}
	}
	return []int{first, last}
}

func Test_645(t *testing.T) {
	NewTestcases(t).
		Add(MakeIntSlice("[2,3]"), MakeIntSlice("[1,2,2,4]")).
		Add(MakeIntSlice("[1,2]"), MakeIntSlice("[1,1]")).
		Add(MakeIntSlice("[1,3]"), MakeIntSlice("[1,1,2]")).
		Add(MakeIntSlice("[2,4]"), MakeIntSlice("[1,3,2,2]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := findErrorNums(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
