package leetcode_test

import (
	"math"
	"math/bits"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "342"
  title: "power-of-four"
  lang: "golang"
  type: "large"
  input:
    - "16\n-128\n5\n1\n2\n64\n-64\n4\n256\n2147483647\n-2147483648\n2147483646\n-2147483647"
*/

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if bits.OnesCount(uint(n)) != 1 {
		return false
	}
	return n&0x55555555 > 0
}

func Test_342(t *testing.T) {
	NewTestcases(t).
		Add(true, 16).
		Add(false, 5).
		Add(true, 1).
		Add(false, 2).
		Add(true, 64).
		Add(false, -64).
		Add(true, 4).
		Add(true, 256).
		Add(false, -128).
		Add(false, math.MaxInt32).
		Add(false, math.MinInt32).
		Add(false, math.MaxInt32-1).
		Add(false, math.MinInt32+1).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isPowerOfFour(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
