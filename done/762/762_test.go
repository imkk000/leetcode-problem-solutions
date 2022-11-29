package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "762"
  title: find-anagram-mappings
  lang: golang
  type: large
  inputString: |-
    [12,28,46,32,50]
    [50,12,32,46,28]
    [84,46]
    [84,46]
*/

func anagramMappings(nums1 []int, nums2 []int) []int {
	for i, n := range nums1 {
		for j := range nums2 {
			if n == nums2[j] {
				nums1[i] = j
			}
		}
	}
	return nums1
}

func Test_762(t *testing.T) {
	type Data struct {
		nums1 []int
		nums2 []int
	}
	NewTestcases(t).
		Add(MakeIntSlice("1,4,3,2,0"), Data{
			nums1: MakeIntSlice("[12,28,46,32,50]"),
			nums2: MakeIntSlice("[50,12,32,46,28]"),
		}).
		Add(MakeIntSlice("[0,1]"), Data{
			nums1: MakeIntSlice("[84,46]"),
			nums2: MakeIntSlice("[84,46]"),
		}).
		Add(MakeIntSlice("[0]"), Data{
			nums1: MakeIntSlice("[0]"),
			nums2: MakeIntSlice("[0]"),
		}).
		Add(MakeIntSlice("[5,2,4,5,4,5]"), Data{
			nums1: MakeIntSlice("[84,8,0,84,0,84]"),
			nums2: MakeIntSlice("[84,84,8,0,0,84]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			nums1, nums2 := input.nums1, input.nums2
			actual := anagramMappings(nums1, nums2)

			a.Equal(td.Expectation, actual)
		})
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1 := []int{84, 8, 0, 84, 0, 84}
		nums2 := []int{84, 84, 8, 0, 0, 84}
		anagramMappings(nums1, nums2)
	}
}
