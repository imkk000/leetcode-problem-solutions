package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1547"
  title: destination-city
  lang: golang
  type: large
  inputString: |-
    [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
    [["B","C"],["D","B"],["C","A"]]
    [["A","Z"]]
    [["A","Z"],["B","A"]]
    [["qMTSlfgZlC","ePvzZaqLXj"],["xKhZXfuBeC","TtnllZpKKg"],["ePvzZaqLXj","sxrvXFcqgG"],["sxrvXFcqgG","xKhZXfuBeC"],["TtnllZpKKg","OAxMijOZgW"]]
*/

func destCity(paths [][]string) string {
	paths[0][0] = paths[0][1]
	m := make(map[string]string)
	for i := range paths[1:] {
		m[paths[i+1][0]] = paths[i+1][1]
	}
	for v := m[paths[0][0]]; v != ""; v = m[v] {
		paths[0][0] = v
	}
	return paths[0][0]
}

func Test_1547(t *testing.T) {
	NewTestcases(t).
		AddExpectation("Sao Paulo").
		AddInput([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}).
		Add("A", [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}).
		Add("Z", [][]string{{"A", "Z"}}).
		Add("Z", [][]string{{"A", "Z"}, {"B", "A"}}).
		AddExpectation("OAxMijOZgW").
		AddInput([][]string{{"qMTSlfgZlC", "ePvzZaqLXj"}, {"xKhZXfuBeC", "TtnllZpKKg"},
			{"ePvzZaqLXj", "sxrvXFcqgG"}, {"sxrvXFcqgG", "xKhZXfuBeC"}, {"TtnllZpKKg", "OAxMijOZgW"}},
		).
		Each(func(a *assert.Assertions, td TestData) {
			actual := destCity(td.Input.([][]string))

			a.Equal(td.Expectation, actual)
		})
}
