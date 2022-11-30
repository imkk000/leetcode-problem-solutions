package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1584"
  title: average-salary-excluding-the-minimum-and-maximum-salary
  lang: golang
  type: large
  inputString: |-
    [4000,3000,1000,2000]
    [1000,2000,3000]
*/

func average(salary []int) float64 {
	sum := salary[0]
	min, max := sum, sum
	l := len(salary)
	for i := 1; i < l; i++ {
		sum += salary[i]
		if min > salary[i] {
			min = salary[i]
		}
		if max < salary[i] {
			max = salary[i]
		}
	}
	return float64(sum-max-min) / float64(l-2)
}

func Test_1584(t *testing.T) {
	NewTestcases(t).
		Add(2500.00000, MakeIntSlice("[4000,3000,1000,2000]")).
		Add(2000.00000, MakeIntSlice("[1000,2000,3000]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := average(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
