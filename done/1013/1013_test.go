package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1013"
  title: fibonacci-number
  lang: golang
  type: large
  inputString: |-
    0
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
    30
*/

func fib(n int) int {
	d := [2]int{0, 1}
	if n <= 1 {
		return d[n]
	}
	for i := 2; i <= n; i++ {
		d[1], d[0] = d[1]+d[0], d[1]
	}
	return d[1]
}

func Test_1013(t *testing.T) {
	NewTestcases(t).
		Add(0, 0).
		Add(1, 1).
		Add(1, 2).
		Add(2, 3).
		Add(3, 4).
		Add(5, 5).
		Add(832040, 30).
		Each(func(a *assert.Assertions, td TestData) {
			actual := fib(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
