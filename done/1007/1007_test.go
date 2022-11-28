package leetcode_test

import (
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1007"
  title: numbers-with-same-consecutive-differences
  lang: golang
  type: large
  inputString: |-
    3
    7
    2
    1
*/

func numsSameConsecDiff(n int, k int) (v []int) {
	var result []int
	var recursion func(i, l, n, k, ans int)
	recursion = func(i, l, n, k, ans int) {
		if i >= l-1 {
			result = append(result, ans)
			return
		}
		if n+k <= 9 {
			recursion(i+1, l, n+k, k, (ans*10)+(n+k))
		}
		if n-k >= 0 && k > 0 {
			recursion(i+1, l, n-k, k, (ans*10)+(n-k))
		}
	}

	for i := 1; i <= 9; i++ {
		recursion(0, n, i, k, i)
	}
	return result
}

func Test_1007(t *testing.T) {
	type Data struct {
		n, k int
	}
	NewTestcases(t).
		Add(MakeIntSlice("[181,292,707,818,929]"), Data{n: 3, k: 7}).
		Add(MakeIntSlice("[10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]"), Data{n: 2, k: 1}).
		Add(MakeIntSlice("[101,121,123,210,212,232,234,321,323,343,345,432,434,454,456,543,545,565,567,654,656,676,678,765,767,787,789,876,878,898,987,989]"), Data{n: 3, k: 1}).
		Add(MakeIntSlice("[111111111,222222222,333333333,444444444,555555555,666666666,777777777,888888888,999999999]"), Data{n: 9, k: 0}).
		Add(MakeIntSlice("[151515151,151515159,151515951,151515959,151595151,151595159,151595951,151595959,159515151,159515159,159515951,159515959,159595151,159595159,159595951,159595959,262626262,373737373,404040404,404040484,404048404,404048484,404840404,404840484,404848404,404848484,484040404,484040484,484048404,484048484,484840404,484840484,484848404,484848484,515151515,515151595,515159515,515159595,515951515,515951595,515959515,515959595,595151515,595151595,595159515,595159595,595951515,595951595,595959515,595959595,626262626,737373737,840404040,840404048,840404840,840404848,840484040,840484048,840484840,840484848,848404040,848404048,848404840,848404848,848484040,848484048,848484840,848484848,951515151,951515159,951515951,951515959,951595151,951595159,951595951,951595959,959515151,959515159,959515951,959515959,959595151,959595159,959595951,959595959]"), Data{n: 9, k: 4}).
		Add(MakeIntSlice("[161616161,272727272,383838383,494949494,505050505,616161616,727272727,838383838,949494949]"), Data{n: 9, k: 5}).
		Add(MakeIntSlice("[171717171,282828282,393939393,606060606,717171717,828282828,939393939]"), Data{n: 9, k: 6}).
		Add(MakeIntSlice("[181818181,292929292,707070707,818181818,929292929]"), Data{n: 9, k: 7}).
		Add(MakeIntSlice("[191919191,808080808,919191919]"), Data{n: 9, k: 8}).
		Add(MakeIntSlice("[909090909]"), Data{n: 9, k: 9}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			n, k := input.n, input.k
			actual := numsSameConsecDiff(n, k)

			sort.Ints(td.Expectation.([]int))
			sort.Ints(actual)
			a.Equal(td.Expectation, actual)
		})
}
