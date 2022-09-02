package leetcode_test

import (
	"math/bits"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "767"
  title: "prime-number-of-set-bits-in-binary-representation"
  lang: "golang"
  type: "large"
  inputString: "990000\n1000000"
  input:
    - "6\n10\n10\n15"
*/

var tb767 = []int{1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0}

func countPrimeSetBits(left, right int) (c int) {
	for left <= right {
		c += tb767[bits.OnesCount(uint(left))]
		left++
	}
	return
}

func Test_767(t *testing.T) {
	NewTestcases(t).
		Add(6, []int{1, 10}).
		Add(4, []int{6, 10}).
		Add(5, []int{10, 15}).
		Add(3754, []int{1e6 - 1e4, 1e6}).
		Add(65, []int{1, 100}).
		Add(3807, []int{524287 - 1e4, 524287}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.([]int)
			left, right := input[0], input[1]
			actual := countPrimeSetBits(left, right)

			a.Equal(td.Expectation, actual)
		})
}
