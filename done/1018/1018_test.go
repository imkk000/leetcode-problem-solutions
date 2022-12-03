package leetcode_test

import (
	_ "embed"
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1018"
  title: largest-perimeter-triangle
  lang: golang
  type: large
  inputString: |-
    [2,1,2]
    [1,2,1,10]
*/

func largestPerimeter(v []int) int {
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})
	l := len(v)
	for i := 0; i < l-2; i++ {
		if v[i] < v[i+1]+v[i+2] {
			return v[i] + v[i+1] + v[i+2]
		}
	}
	return 0
}

func Test_1018(t *testing.T) {
	NewTestcases(t).
		Add(5, MakeIntSlice("[2,1,2]")).
		Add(0, MakeIntSlice("[1,2,1,10]")).
		Add(3700, MakeIntSlice(input)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := largestPerimeter(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}

// go:exclude
//
//go:embed 1018_1_in
var input string
