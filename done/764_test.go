package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "764"
  title: n-ary-tree-level-order-traversal
  lang: golang
  type: large
  inputString: |-
    [1,null,3,2,4,null,5,6]
    [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    []
    [1]
*/

// go:exclude
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	var recursion func(result *[][]int, root *Node, i int)
	recursion = func(result *[][]int, root *Node, i int) {
		if i == len(*result) {
			*result = append(*result, make([]int, 0))
		}
		(*result)[i] = append((*result)[i], root.Val)
		for _, node := range root.Children {
			recursion(result, node, i+1)
		}
	}
	recursion(&result, root, 0)
	return result
}

func Test_764(t *testing.T) {
	NewTestcases(t).
		AddExpectation(Make2DMatrixInt("[[1],[3,2,4],[5,6]]")).
		AddInput(&Node{
			Val: 1,
			Children: []*Node{{
				Val:      3,
				Children: []*Node{{Val: 5}, {Val: 6}},
			}, {Val: 2}, {Val: 4}},
		}).
		Each(func(a *assert.Assertions, td TestData) {
			actual := levelOrder(td.Input.(*Node))

			a.Equal(td.Expectation, actual)
		})
}
