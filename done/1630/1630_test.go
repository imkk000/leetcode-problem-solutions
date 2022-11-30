package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1630"
  title: count-odd-numbers-in-an-interval-range
  lang: golang
  type: large
  inputString: |-
    3
    7
    8
    10
*/

func countOdds(low int, high int) int {
	if low&1 == 0 {
		low++
	}
	if high&1 == 0 {
		high--
	}
	if high = ((high - low) / 2) + 1; high > 0 {
		return high
	}
	return 0
}

func Test_1630(t *testing.T) {
	type Data struct {
		low, high int
	}
	NewTestcases(t).
		Add(3, Data{low: 3, high: 7}).
		Add(1, Data{low: 8, high: 10}).
		Add(3, Data{low: 0, high: 5}).
		Add(3, Data{low: 1, high: 6}).
		Add(1, Data{low: 2, high: 3}).
		Add(1, Data{low: 2, high: 4}).
		Add(1, Data{low: 1, high: 1}).
		Add(0, Data{low: 2, high: 2}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			low, high := input.low, input.high
			actual := countOdds(low, high)

			a.Equal(td.Expectation, actual)
		})
}
