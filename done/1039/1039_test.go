package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1039"
  title: find-the-town-judge
  lang: golang
  type: large
  inputString: |-
    2
    [[1,2]]
    3
    [[1,3],[2,3]]
    3
    [[1,3],[2,3],[3,1]]
    3
    [[1,3],[3,1]]
    2
    [[2,1]]
    1
    []
    2
    [[1,2],[2,1]]
    3
    [[1,2],[1,3]]
    3
    [[2,1],[3,2]]
*/

func findJudge(n int, trust [][]int) (trustPerson int) {
	if n <= 1 {
		return 1
	}
	trustPoint := make([]int, n+1)
	hasTrustSomeone := make([]bool, n+1)
	for _, t := range trust {
		hasTrustSomeone[t[0]] = true
		if trustPoint[t[1]]++; trustPoint[t[1]] == n-1 {
			trustPerson = t[1]
		}
	}
	if trustPerson == 0 || hasTrustSomeone[trustPerson] {
		return -1
	}
	return
}

func Test_1039(t *testing.T) {
	type Data struct {
		n     int
		trust [][]int
	}
	NewTestcases(t).
		Add(-1, Data{
			n:     4,
			trust: Make2DMatrixInt("[[1,3],[1,4],[2,3]]"),
		}).
		Add(-1, Data{
			n:     3,
			trust: Make2DMatrixInt("[[2,1],[3,2]]"),
		}).
		Add(2, Data{
			n:     2,
			trust: Make2DMatrixInt("[[1,2]]"),
		}).
		Add(3, Data{
			n:     3,
			trust: Make2DMatrixInt("[[1,3],[2,3]]"),
		}).
		Add(-1, Data{
			n:     3,
			trust: Make2DMatrixInt("[[1,3],[2,3],[3,1]]"),
		}).
		Add(1, Data{
			n:     1,
			trust: Make2DMatrixInt("[]"),
		}).
		Add(-1, Data{
			n:     2,
			trust: Make2DMatrixInt("[[1,2],[2,1]]"),
		}).
		Add(-1, Data{
			n:     3,
			trust: Make2DMatrixInt("[[1,2],[1,3]]"),
		}).
		Add(3, Data{
			n:     4,
			trust: Make2DMatrixInt("[[1,3],[1,4],[2,3],[2,4],[4,3]]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			n, trust := input.n, input.trust
			actual := findJudge(n, trust)

			a.Equal(td.Expectation, actual)
		})
}
