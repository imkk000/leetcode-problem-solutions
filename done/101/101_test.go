package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "101"
  title: symmetric-tree
  lang: golang
  type: large
  inputString: |-
    [1,2,2,3,4,4,3]
    [1,2,2,null,3,null,3]
*/

func isSymmetric(root *TreeNode) bool {
	return recursion(root.Left, root.Right)
}

func recursion(left, right *TreeNode) bool {
	if left == right && left == nil {
		return true
	}
	if left == nil && right != nil ||
		left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return recursion(left.Left, right.Right) &&
		recursion(left.Right, right.Left)
}

func Test_101(t *testing.T) {
	NewTestcases(t).
		Add(true, MakeTreeNode("[1,2,2,3,4,4,3]")).
		Add(true, MakeTreeNode("[1]")).
		Add(true, MakeTreeNode("[1,2,2]")).
		Add(false, MakeTreeNode("[1,2,3]")).
		Add(false, MakeTreeNode("[1,2,2,null,3,null,3]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := isSymmetric(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
