package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1236"
  title: n-th-tribonacci-number
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
    25
    37
*/

func tribonacci(n int) int {
	d := [3]int{0, 1, 1}
	if n <= 2 {
		return d[n]
	}
	for i := 3; i <= n; i++ {
		d[2], d[1], d[0] = d[2]+d[1]+d[0], d[2], d[1]
	}
	return d[2]
}

func Test_1236(t *testing.T) {
	NewTestcases(t).
		Add(0, 0).
		Add(1, 1).
		Add(1, 2).
		Add(2, 3).
		Add(4, 4).
		Add(7, 5).
		Add(13, 6).
		Add(1389537, 25).
		Add(2082876103, 37).
		Each(func(a *assert.Assertions, td TestData) {
			actual := tribonacci(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
