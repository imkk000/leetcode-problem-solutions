package leetcode_test

import (
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1354"
  title: find-players-with-zero-or-one-losses
  lang: golang
  type: large
  inputString: |-
    [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
    [[2,3],[1,3],[5,4],[6,4]]
*/

func findWinners(matches [][]int) [][]int {
	s := make(map[int][]int)
	for _, m := range matches {
		if _, exists := s[m[0]]; !exists {
			s[m[0]] = make([]int, 2)
		}
		if _, exists := s[m[1]]; !exists {
			s[m[1]] = make([]int, 2)
		}
		s[m[0]][0]++
		s[m[1]][1]++
	}
	matches = [][]int{{}, {}}
	for id, v := range s {
		if v[1] > 1 {
			continue
		}
		matches[v[1]] = append(matches[v[1]], id)
	}
	sort.Ints(matches[0])
	sort.Ints(matches[1])
	return matches
}

func Test_1354(t *testing.T) {
	NewTestcases(t).
		AddExpectation(Make2DMatrixInt("[[1,2,10],[4,5,7,8]]")).
		AddInput(Make2DMatrixInt("[[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]")).
		Add(Make2DMatrixInt("[[1,2,5,6],[]]"), Make2DMatrixInt("[[2,3],[1,3],[5,4],[6,4]]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := findWinners(td.Input.([][]int))

			a.Equal(td.Expectation, actual)
		})
}
