package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1149"
  title: intersection-of-three-sorted-arrays
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4,5]
    [1,2,5,7,9]
    [1,3,4,5,8]
    [197,418,523,876,1356]
    [501,880,1593,1710,1870]
    [521,682,1337,1395,1764]
*/

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) (s []int) {
	var v [2001]int
	for _, n := range arr1 {
		v[n]++
	}
	for _, n := range arr2 {
		v[n]++
	}
	for _, n := range arr3 {
		if v[n] == 2 {
			s = append(s, n)
		}
	}
	return
}

func Test_1149(t *testing.T) {
	type Data struct {
		arr1, arr2, arr3 []int
	}
	var intSliceNil []int
	NewTestcases(t).
		AddExpectation(MakeIntSlice("[1,5]")).
		AddInput(Data{
			arr1: MakeIntSlice("[1,2,3,4,5]"),
			arr2: MakeIntSlice("[1,2,5,7,9]"),
			arr3: MakeIntSlice("[1,3,4,5,8]"),
		}).
		AddExpectation(intSliceNil).
		AddInput(Data{
			arr1: MakeIntSlice("[197,418,523,876,1356]"),
			arr2: MakeIntSlice("[501,880,1593,1710,1870]"),
			arr3: MakeIntSlice("[521,682,1337,1395,1764]"),
		}).
		AddExpectation(MakeIntSlice("1")).
		AddInput(Data{
			arr1: MakeIntSlice("[1,2,3,4,5]"),
			arr2: MakeIntSlice("[1,2]"),
			arr3: MakeIntSlice("[1,3,4,5]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			arr1, arr2, arr3 := input.arr1, input.arr2, input.arr3
			actual := arraysIntersection(arr1, arr2, arr3)

			a.Equal(td.Expectation, actual)
		})
}
