package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2491"
  title: smallest-even-multiple
  lang: golang
  type: large
  inputString: |-
    1
    5
    6
    150
*/

func smallestEvenMultiple(n int) int {
	if n&1 == 1 {
		return n * 2
	}
	return n
}

func Test_2491(t *testing.T) {
	NewTestcases(t).
		Add(10, 5).
		Add(6, 6).
		Each(func(a *assert.Assertions, td TestData) {
			actual := smallestEvenMultiple(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
