package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "832"
  title: binary-tree-pruning
  lang: golang
  type: large
  inputString: |-
    [1,null,0,0,1]
    [1,0,1,0,0,0,1]
    [1,1,0,1,1,0,1,0]
    [0]
    [1]
    [0,0,0,0,0,0,0]
*/

func pruneTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		root = nil
	}
	return root
}

func Test_832(t *testing.T) {
	var nodeNil *TreeNode
	NewTestcases(t).
		Add(nodeNil, MakeTreeNode("[0]")).
		Add(nodeNil, MakeTreeNode("[0,0,0,0,0,0]")).
		Add(MakeTreeNode("[1,1,0,1,1,null,1]"), MakeTreeNode("[1,1,0,1,1,0,1,0]")).
		Add(MakeTreeNode("[1,null,1,null,1]"), MakeTreeNode("[1,0,1,0,0,0,1]")).
		Add(MakeTreeNode("[1,null,0,null,1]"), MakeTreeNode("[1,null,0,0,1]")).
		Add(MakeTreeNode("[1]"), MakeTreeNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := pruneTree(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
