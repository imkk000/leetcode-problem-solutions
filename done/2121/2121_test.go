package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"

	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2121"
  title: find-if-path-exists-in-graph
  lang: golang
  type: large
  inputString: |-
    3
    [[0,1],[1,2],[2,0]]
    0
    2
    6
    [[0,1],[0,2],[3,5],[5,4],[4,3]]
    0
    5
    8
    [[1,7],[7,2],[2,6],[6,5],[5,4],[4,3]]
    1
    3
*/

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	visited := make([]bool, n)
	return walk(g, visited, source, destination)
}

func walk(g [][]int, visited []bool, src, dest int) (v bool) {
	if src == dest {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true
	for _, next := range g[src] {
		if visited[next] {
			continue
		}
		v = v || walk(g, visited, next, dest)
	}
	return
}

func Test_2121(t *testing.T) {
	type Data struct {
		n, src, dest int
		edges        [][]int
	}
	NewTestcases(t).
		Add(true, Data{
			n:     3,
			src:   0,
			dest:  2,
			edges: Make2DMatrixInt("[[0,1],[1,2],[2,0]]"),
		}).
		Add(false, Data{
			n:     6,
			src:   0,
			dest:  5,
			edges: Make2DMatrixInt("[[0,1],[0,2],[3,5],[5,4],[4,3]]"),
		}).
		Add(true, Data{
			n:     2,
			src:   0,
			dest:  1,
			edges: Make2DMatrixInt("[[0,1]]"),
		}).
		Add(true, Data{}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			n, edges, source, destination := input.n, input.edges, input.src, input.dest
			actual := validPath(n, edges, source, destination)

			a.Equal(td.Expectation, actual)
		})
}
