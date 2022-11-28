package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "326"
  title: power-of-three
  lang: golang
  type: large
  inputString: |-
    27
    0
    9
    -9
    2147483647
    -2147483648
*/

func isPowerOfThree(n int) bool {
	return n >= 1 && 1162261467%n == 0
}

func Test_326(t *testing.T) {
	NewTestcases(t).
		Add(true, 1).
		Add(true, 3).
		Add(true, 9).
		Add(true, 27).
		Add(true, 81).
		Add(false, 242).
		Add(true, 243).
		Add(false, 244).
		Add(true, 129140163).
		Add(false, 19682).
		Add(false, 1594322).
		Add(false, 0).
		Add(false, 2).
		Add(false, 4).
		Add(false, -3).
		Add(false, -9).
		Add(false, math.MaxInt32).
		Add(false, math.MinInt32).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isPowerOfThree(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
