package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1406"
  title: subtract-the-product-and-sum-of-digits-of-an-integer
  lang: golang
  type: large
  inputString: |-
    234
    4421
    12
*/

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for n > 0 {
		p *= n % 10
		s += n % 10
		n /= 10
	}
	return p - s
}

func Test_1406(t *testing.T) {
	NewTestcases(t).
		Add(15, 234).
		Add(21, 4421).
		Add(-1, 12).
		Each(func(a *assert.Assertions, td TestData) {
			actual := subtractProductAndSum(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
