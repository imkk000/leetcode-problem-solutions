package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1800"
  title: concatenation-of-consecutive-binary-numbers
  lang: golang
  type: large
  inputString: |-
    1
    2
    3
    4
    12
    100000
*/

func concatenatedBinary(n int) int {
	const mod1e9p7 = 1_000_000_007

	s, b, l := 1, 2, 1
	for i := 2; i <= n; i++ {
		if b == i {
			b <<= 1
			l++
		}
		s = (s<<l)%mod1e9p7 + i
	}
	return s % mod1e9p7
}

func Test_1800(t *testing.T) {
	NewTestcases(t).
		Add(1, 1).
		Add(6, 2).
		Add(27, 3).
		Add(220, 4).
		Add(505379714, 12).
		Add(880407397, 99999).
		Add(757631812, 100000).
		Each(func(a *assert.Assertions, td TestData) {
			actual := concatenatedBinary(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
