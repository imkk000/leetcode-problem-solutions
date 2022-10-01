package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2502"
  title: sort-the-people
  lang: golang
  type: large
  inputString: |-
    ["Mary","John","Emma"]
    [180,165,170]
    ["Alice","Bob","Bob"]
    [155,185,150]
*/

func sortPeople(names []string, heights []int) []string {
	l := len(names)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if heights[i] < heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
				names[i], names[j] = names[j], names[i]
			}
		}
	}
	return names
}

func Test_2502(t *testing.T) {
	type Data struct {
		names   []string
		heights []int
	}
	NewTestcases(t).
		Add(MakeStringSlice(`["Mary","Emma","John"]`), Data{
			names:   MakeStringSlice(`["Mary","John","Emma"]`),
			heights: MakeIntSlice("[180,165,170]"),
		}).
		Add(MakeStringSlice(`["Bob","Alice","Bob"]`), Data{
			names:   MakeStringSlice(`["Alice","Bob","Bob"]`),
			heights: MakeIntSlice("[155,185,150]"),
		}).
		AddExpectation(MakeStringSlice(`["EPCFFt","RPJOFYZUBFSIYp","VOYGWWNCf","Vk","Sgizfdfrims","IEO","QTASHKQ","WSpmqvb"]`)).
		AddInput(Data{
			names:   MakeStringSlice(`["IEO","Sgizfdfrims","QTASHKQ","Vk","RPJOFYZUBFSIYp","EPCFFt","VOYGWWNCf","WSpmqvb"]`),
			heights: MakeIntSlice("[17233,32521,14087,42738,46669,65662,43204,8224]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := sortPeople(input.names, input.heights)

			a.Equal(td.Expectation, actual)
		})
}
