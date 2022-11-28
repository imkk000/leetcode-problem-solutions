package leetcode_test

import (
	_ "embed"
	"math/bits"
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1458"
  title: sort-integers-by-the-number-of-1-bits
  lang: golang
  type: large
  inputString: |-
    [0,1,2,3,4,5,6,7,8]
    [1024,512,256,128,64,32,16,8,4,2,1]
*/

type SortInts []int

func (s SortInts) Len() int {
	return len(s)
}

func (s SortInts) Less(i, j int) bool {
	a := bits.OnesCount32(uint32(s[i]))
	b := bits.OnesCount32(uint32(s[j]))
	if a != b {
		return a < b
	}
	return s[i] < s[j]
}

func (s SortInts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortByBits(arr []int) []int {
	sort.Sort(SortInts(arr))
	return arr
}

func Test_1458(t *testing.T) {
	NewTestcases(t).
		Add(MakeIntSlice(exp1458), newData(1e4)).
		Add(MakeIntSlice("[1,2,4,8,16,32,64,128,256,512,1024]"), MakeIntSlice("[1024,512,256,128,64,32,16,8,4,2,1]")).
		Add(MakeIntSlice("[0,1,2,4,8,3,5,6,7]"), MakeIntSlice("[0,1,2,3,4,5,6,7,8]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := sortByBits(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}

// go:exclude
var (
	newData = func(n int) []int {
		v := make([]int, n+1)
		for i := 0; i <= n; i++ {
			v[i] = i
		}
		return v
	}
	benchData = newData(1e4)
)

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByBits(benchData)
	}
}

// go:exclude
//
//go:embed 1458_1_exp
var exp1458 string
