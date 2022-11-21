package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "222"
  title: count-complete-tree-nodes
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4,5,6]
    []
    [1]
*/

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var recursion func(root *TreeNode, c *int)
	recursion = func(root *TreeNode, c *int) {
		*c++
		if root.Left != nil {
			recursion(root.Left, c)
		}
		if root.Right != nil {
			recursion(root.Right, c)
		}
	}

	var c int
	recursion(root, &c)
	return c
}

func Test_222(t *testing.T) {
	NewTestcases(t).
		Add(6, MakeTreeNode("[1,2,3,4,5,6]")).
		Add(0, MakeTreeNode("[]")).
		Add(1, MakeTreeNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := countNodes(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
