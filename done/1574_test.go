package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1574"
  title: maximum-product-of-two-elements-in-an-array
  lang: golang
  type: large
  inputString: |-
    [3,4,5,2]
    [1,5,4,5]
    [3,7]
    [1,1]
    [1,2]
    [2,2]
*/

func maxProduct(nums []int) int {
	var m1, m2 int
	for _, v := range nums {
		if m1 < v {
			m1, m2 = v, m1
			continue
		}
		if m2 < v && m1 > m2 {
			m2 = v
		}
	}
	return (m1 - 1) * (m2 - 1)
}

func Test_1574(t *testing.T) {
	NewTestcases(t).
		Add(12, MakeIntSlice("[3,4,5,2]")).
		Add(16, MakeIntSlice("[1,5,4,5]")).
		Add(12, MakeIntSlice("[3,7]")).
		Add(0, MakeIntSlice("[1,1]")).
		Add(0, MakeIntSlice("[1,2]")).
		Add(1, MakeIntSlice("[2,2]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := maxProduct(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
