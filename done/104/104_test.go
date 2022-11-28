package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "104"
  title: maximum-depth-of-binary-tree
  lang: golang
  type: large
  inputString: |-
    [3,9,20,null,null,15,7]
    [1,null,2]
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var mx int
	var recursion func(root *TreeNode, d int)
	recursion = func(root *TreeNode, d int) {
		if mx < d {
			mx = d
		}
		if root.Left != nil {
			recursion(root.Left, d+1)
		}
		if root.Right != nil {
			recursion(root.Right, d+1)
		}
	}
	recursion(root, 1)
	return mx
}

func Test_104(t *testing.T) {
	NewTestcases(t).
		Add(3, MakeTreeNode("[3,9,20,null,null,15,7]")).
		Add(2, MakeTreeNode("[1,null,2]")).
		Add(0, MakeTreeNode("[]")).
		Add(1, MakeTreeNode("[2]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := maxDepth(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
