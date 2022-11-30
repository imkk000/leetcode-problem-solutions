package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1319"
  title: unique-number-of-occurrences
  lang: golang
  type: large
  inputString: |-
    [1,2,2,1,1,3]
    [1,2]
    [-3,0,1,-3,1,1,1,-3,10,0]
*/

func uniqueOccurrences(nums []int) bool {
	var v [2001]int
	var s []int
	for _, n := range nums {
		v[n+1000]++
		if v[n+1000] == 1 {
			s = append(s, n)
		}
	}
	var d [1001]bool
	for _, n := range s {
		if d[v[n+1000]] {
			return false
		}
		d[v[n+1000]] = true
	}
	return true
}

func Test_1319(t *testing.T) {
	NewTestcases(t).
		Add(true, MakeIntSlice("[1,2,2,1,1,3]")).
		Add(false, MakeIntSlice("[1,2]")).
		Add(true, MakeIntSlice("[-3,0,1,-3,1,1,1,-3,10,0]")).
		Add(false, MakeIntSlice("[-1000,0,1000]")).
		Add(true, MakeIntSlice("[0]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := uniqueOccurrences(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
