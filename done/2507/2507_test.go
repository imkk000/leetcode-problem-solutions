package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2507"
  title: number-of-common-factors
  lang: golang
  type: large
  inputString: |-
    12
    6
    25
    30
*/

func commonFactors(a int, b int) (c int) {
	if a > b {
		a, b = b, a
	}
	for i := 1; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			c++
		}
	}
	return
}

func Test_2507(t *testing.T) {
	type Data struct {
		a, b int
	}
	NewTestcases(t).
		AddExpectation(4).
		AddInput(Data{
			a: 12,
			b: 6,
		}).
		AddExpectation(2).
		AddInput(Data{
			a: 25,
			b: 30,
		}).
		AddExpectation(1).
		AddInput(Data{
			a: 1,
			b: 1000,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := commonFactors(input.a, input.b)

			a.Equal(td.Expectation, actual)
		})
}
