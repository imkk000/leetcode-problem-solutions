package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2571"
  title: find-the-pivot-integer
  lang: golang
  type: large
  inputString: |-
    1
    2
    3
    4
    5
    6
    7
    8
    9
    10
    999
    1000
*/

func pivotInteger(n int) int {
	var l, r int
	for i := 1; i <= n; i++ {
		r += i
	}
	for i := 1; i <= n; i++ {
		l += i
		r -= (i - 1)
		if l == r {
			return i
		}
	}
	return -1
}

func Test_2571(t *testing.T) {
	NewTestcases(t).
		Add(1, 1).
		Add(-1, 2).
		Add(-1, 3).
		Add(-1, 4).
		Add(-1, 5).
		Add(-1, 6).
		Add(-1, 7).
		Add(6, 8).
		Add(-1, 9).
		Add(-1, 10).
		Add(-1, 999).
		Add(-1, 1000).
		Each(func(a *assert.Assertions, td TestData) {
			actual := pivotInteger(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
