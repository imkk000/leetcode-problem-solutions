package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1090"
  title: armstrong-number
  lang: golang
  type: large
  inputString: |-
    153
    123
    1
*/

func isArmstrong(n int) bool {
	l := math.Floor(math.Log10(float64(n))) + 1
	m := n
	var s int
	for m > 0 {
		s += int(math.Pow(float64(m%10), l))
		m /= 10
	}
	return s == n
}

func Test_1090(t *testing.T) {
	NewTestcases(t).
		Add(true, 153).
		Add(false, 123).
		Add(true, 1).
		Add(true, 2).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isArmstrong(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
