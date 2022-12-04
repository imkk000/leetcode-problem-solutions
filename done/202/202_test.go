package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "202"
  title: happy-number
  lang: golang
  type: large
  inputString: |-
    19
    2
    1
    2147483647
*/

func isHappy(n int) bool {
	var b [1001]bool
	for n > 1 {
		if n = sumDigits(n); b[n] {
			return false
		}
		b[n] = true
	}
	return true
}

func sumDigits(n int) (sum int) {
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return
}

func Test_202(t *testing.T) {
	NewTestcases(t).
		Add(true, 19).
		Add(false, 2).
		Add(true, 1).
		Add(false, math.MaxInt32).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isHappy(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
