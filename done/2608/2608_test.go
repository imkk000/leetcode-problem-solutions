package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2608"
  title: count-the-digits-that-divide-a-number
  lang: golang
  type: large
  inputString: |-
    7
    121
    1248
    39
    1
    99999999
*/

func countDigits(n int) (s int) {
	num := n
	for n > 0 {
		if num%(n%10) == 0 {
			s++
		}
		n /= 10
	}
	return
}

func Test_2608(t *testing.T) {
	NewTestcases(t).
		Add(1, 7).
		Add(2, 121).
		Add(4, 1248).
		Add(1, 39).
		Add(1, 1).
		Add(8, 99999999).
		Each(func(a *assert.Assertions, td TestData) {
			actual := countDigits(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
