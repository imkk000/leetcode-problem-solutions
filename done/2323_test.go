package leetcode_test

import (
	"math/bits"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2323"
  title: minimum-bit-flips-to-convert-number
  lang: golang
  type: large
  inputString: |-
    10
    7
    3
    4
    0
    1000000000
    1000000000
    0
    1000000000
    1
    1
    1
*/

func minBitFlips(start, goal int) int {
	return bits.OnesCount32(uint32(start ^ goal))
}

func Test_2323(t *testing.T) {
	NewTestcases(t).
		Add(3, []int{10, 7}).
		Add(3, []int{3, 4}).
		Add(13, []int{0, 1e9}).
		Add(13, []int{1e9, 0}).
		Add(14, []int{1, 1e9}).
		Add(0, []int{1, 1}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.([]int)
			start, goal := input[0], input[1]
			actual := minBitFlips(start, goal)

			a.Equal(td.Expectation, actual)
		})
}
