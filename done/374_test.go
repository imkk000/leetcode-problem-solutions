package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "374"
  title: guess-number-higher-or-lower
  lang: golang
  type: large
  inputString: |-
    10
    6
    1
    1
    2
    1
    1000
    50
*/

func guessNumber(high int) int {
	low := 1
	for low <= high {
		mid := ((high - low) / 2) + low
		g := guess374(mid)
		if g > 0 {
			low = mid + 1
		} else if g < 0 {
			high = mid - 1
		} else {
			return mid
		}
	}
	return high
}

// go:exclude
var pick374 int

// go:exclude
func guess374(n int) int {
	if n > pick374 {
		return -1
	}
	if n < pick374 {
		return 1
	}
	return 0
}

func Test_374(t *testing.T) {
	type Data struct {
		num  int
		pick int
	}
	NewTestcases(t).
		Add(50, Data{
			num:  1000,
			pick: 50,
		}).
		Add(6, Data{
			num:  10,
			pick: 6,
		}).
		Add(1, Data{
			num:  1,
			pick: 1,
		}).
		Add(1, Data{
			num:  2,
			pick: 1,
		}).
		Add(1, Data{
			num:  math.MaxInt32,
			pick: 1,
		}).
		Add(math.MaxInt32/2, Data{
			num:  math.MaxInt32,
			pick: math.MaxInt32 / 2,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			pick374 = input.pick
			actual := guessNumber(input.num)

			a.Equal(td.Expectation, actual)
		})
}
