package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "263"
  title: ugly-number
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
    11
    12
    13
    14
    15
    2147483647
*/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		switch {
		case n%2 == 0:
			n /= 2
		case n%3 == 0:
			n /= 3
		case n%5 == 0:
			n /= 5
		default:
			return false
		}
	}
	return true
}

func Test_263(t *testing.T) {
	NewTestcases(t).
		Add(true, 6).
		Add(false, -6).
		Add(true, 1).
		Add(false, 14).
		Add(false, 28).
		Add(false, math.MaxInt32).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isUgly(td.Input.(int))
			a.Equal(td.Expectation, actual)
		})
}
