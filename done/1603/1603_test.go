package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1603"
  title: running-sum-of-1d-array
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4]
    [1,1,1,1,1]
    [3,1,2,10,1]
*/

func runningSum(n []int) []int {
	for i := range n[1:] {
		n[i+1] += n[i]
	}
	return n
}

func Test_1603(t *testing.T) {
	NewTestcases(t).
		Add(MakeIntSlice("[1,3,6,10]"), MakeIntSlice("[1,2,3,4]")).
		Add(MakeIntSlice("[1,2,3,4,5]"), MakeIntSlice("[1,1,1,1,1]")).
		Add(MakeIntSlice("[3,4,6,16,17]"), MakeIntSlice("[3,1,2,10,1]")).
		Add(MakeIntSlice("[1]"), MakeIntSlice("[1]")).
		Add(MakeIntSlice("[1,2]"), MakeIntSlice("[1,1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := runningSum(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
