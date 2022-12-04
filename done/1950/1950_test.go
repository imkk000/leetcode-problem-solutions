package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1950"
  title: sign-of-the-product-of-an-array
  lang: golang
  type: large
  inputString: |-
    [-1,-2,-3,-4,3,2,1]
    [1,5,0,2,-3]
    [-1,1,-1,1,-1]
*/

func arraySign(v []int) int {
	p := 1
	for i := range v {
		p *= signFunc(v[i])
	}
	return signFunc(p)
}

func signFunc(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Test_1950(t *testing.T) {
	NewTestcases(t).
		Add(1, MakeIntSlice("[-1,-2,-3,-4,3,2,1]")).
		Add(0, MakeIntSlice("[1,5,0,2,-3]")).
		Add(-1, MakeIntSlice("[-1,1,-1,1,-1]")).
		Add(1, MakeIntSlice("[1]")).
		Add(-1, MakeIntSlice("[9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := arraySign(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
