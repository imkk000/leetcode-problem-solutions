package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1833"
  title: find-the-highest-altitude
  lang: golang
  type: large
  inputString: |-
    [-5,1,5,0,-7]
    [-4,-3,-2,-1,4,3,2]
    [1]
    [-1]
*/

func largestAltitude(g []int) (mx int) {
	var s int
	for _, n := range g {
		if s += n; mx < s {
			mx = s
		}
	}
	return
}

func Test_1833(t *testing.T) {
	NewTestcases(t).
		Add(1, MakeIntSlice("[-5,1,5,0,-7]")).
		Add(0, MakeIntSlice("[-4,-3,-2,-1,4,3,2]")).
		Add(1, MakeIntSlice("[1]")).
		Add(0, MakeIntSlice("[-1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := largestAltitude(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
