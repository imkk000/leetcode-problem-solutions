package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "231"
  title: power-of-two
  lang: golang
  type: large
  inputString: |-
    1
    16
    3
    -3
    -9
    2147483647
    -2147483648
*/

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func Test_231(t *testing.T) {
	NewTestcases(t).
		Add(true, 1).
		Add(true, 2).
		Add(true, 4).
		Add(true, 8).
		Add(true, 16).
		Add(true, 32).
		Add(false, 3).
		Add(false, 5).
		Add(false, -2).
		Add(false, -4).
		Add(false, 0).
		Add(false, math.MaxInt32).
		Add(false, math.MinInt32).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isPowerOfTwo(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
