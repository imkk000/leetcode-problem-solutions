package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "933"
  title: increasing-order-search-tree
  lang: golang
  type: large
  inputString: |-
    [5,3,6,2,4,null,8,1,null,null,null,7,9]
    [5,1,7]
*/

func increasingBST(root *TreeNode) *TreeNode {
	var recursion func(root, head *TreeNode) *TreeNode
	recursion = func(root, head *TreeNode) *TreeNode {
		if root == nil {
			return head
		}
		tail := recursion(root.Left, head)
		tail.Right = root
		root.Left = nil
		return recursion(root.Right, tail.Right)
	}

	head := new(TreeNode)
	recursion(root, head)
	return head.Right
}

func Test_933(t *testing.T) {
	NewTestcases(t).
		AddExpectation(MakeTreeNode("[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]")).
		AddInput(MakeTreeNode("[5,3,6,2,4,null,8,1,null,null,null,7,9]")).
		Add(MakeTreeNode("[1,null,5,null,7]"), MakeTreeNode("[5,1,7]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := increasingBST(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
