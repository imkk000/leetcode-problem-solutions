package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "70"
  title: climbing-stairs
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
    45
*/

func climbStairs(n int) int {
	d := [2]int{1, 1}
	for i := 2; i <= n; i++ {
		d[0], d[1] = d[1]+d[0], d[0]
	}
	return d[0]
}

func Test_70(t *testing.T) {
	NewTestcases(t).
		Add(1, 1).
		Add(2, 2).
		Add(3, 3).
		Add(5, 4).
		Add(8, 5).
		Add(13, 6).
		Each(func(a *assert.Assertions, td TestData) {
			actual := climbStairs(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
