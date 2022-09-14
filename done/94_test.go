package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "94"
  title: binary-tree-inorder-traversal
  lang: golang
  type: large
  inputString: |-
    [1,null,2,3]
    []
    [1]
    [9,8,7,null,3,5,7]
    [2,3,null,1]
    [1,2,7,3,4,null,8]
*/

func inorderTraversal(root *TreeNode) (v []int) {
	if root == nil {
		return
	}
	var recursion func(root *TreeNode) []int
	recursion = func(root *TreeNode) (v []int) {
		var left, right []int
		if root.Left != nil {
			left = recursion(root.Left)
		}
		if root.Right != nil {
			right = recursion(root.Right)
		}

		if left != nil {
			v = append(v, left...)
		}
		v = append(v, root.Val)
		if right != nil {
			v = append(v, right...)
		}
		return
	}
	return recursion(root)
}

func Test_94(t *testing.T) {
	NewTestcases(t).
		Add([]int(nil), MakeTreeNode("[]")).
		Add(MakeIntSlice("[1,3,2]"), MakeTreeNode("[2,3,null,1]")).
		Add(MakeIntSlice("[8,3,9,5,7,7]"), MakeTreeNode("[9,8,7,null,3,5,7]")).
		Add(MakeIntSlice("[1,3,2]"), MakeTreeNode("[1,null,2,3]")).
		Add(MakeIntSlice("[1]"), MakeTreeNode("[1]")).
		Add(MakeIntSlice("[3,2,4,1,7,8]"), MakeTreeNode("[1,2,7,3,4,null,8]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := inorderTraversal(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
