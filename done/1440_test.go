package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1440"
  title: convert-integer-to-the-sum-of-two-no-zero-integers
  lang: golang
  type: large
  inputString: |-
    2
    11
*/

func getNoZeroIntegers(n int) []int {
	m := n
	for containsZero(m) || containsZero(n-m) {
		m--
	}
	return []int{n - m, m}
}

func containsZero(n int) bool {
	if n == 0 {
		return true
	}
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}

func Test_1440(t *testing.T) {
	NewTestcases(t).
		Add([]int{1, 1}, 2).
		Add([]int{2, 9}, 11).
		Add([]int{2, 49}, 51).
		Add([]int{11, 99}, 110).
		Add([]int{1, 549}, 550).
		Each(func(a *assert.Assertions, td TestData) {
			actual := getNoZeroIntegers(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
