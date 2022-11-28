package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1027"
  title: sum-of-even-numbers-after-queries
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4]
    [[1,0],[-3,1],[-4,0],[2,3]]
    [1]
    [[4,0]]
*/

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	b := make([]bool, len(queries))
	var s int
	for i, n := range nums {
		if isEven(n) {
			b[i] = true
			s += n
		}
	}

	v := make([]int, len(queries))
	for i, q := range queries {
		val, index := q[0], q[1]
		nums[index] += val
		if b[index] {
			s += val
		}
		if isEven(nums[index]) {
			if !b[index] {
				s += nums[index]
				b[index] = true
			}
		} else {
			if b[index] {
				s -= nums[index]
				b[index] = false
			}
		}
		v[i] = s
	}
	return v
}

func isEven(n int) bool {
	return n%2 == 0
}

func Test_1027(t *testing.T) {
	type Data struct {
		nums    []int
		queries [][]int
	}
	NewTestcases(t).
		Add(MakeIntSlice("[8,6,2,4]"), Data{
			nums:    MakeIntSlice("[1,2,3,4]"),
			queries: Make2DMatrixInt("[[1,0],[-3,1],[-4,0],[2,3]]"),
		}).
		Add(MakeIntSlice("[0]"), Data{
			nums:    MakeIntSlice("[1]"),
			queries: Make2DMatrixInt("[[4,0]]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := sumEvenAfterQueries(input.nums, input.queries)

			a.Equal(td.Expectation, actual)
		})
}
