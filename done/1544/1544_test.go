package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1544"
  title: count-good-nodes-in-binary-tree
  lang: golang
  type: large
  inputString: |-
    [3,1,4,3,null,1,5,2,5]
    [3,1,4,3,null,1,5]
    [3,3,null,4,2]
    [1]
    [-10000]
    [10000]
    [2,null,4,10,8,null,null,4]
    [3,3,null,4,2]
    [9,null,3,6]
*/

func goodNodes(root *TreeNode) int {
	var recursion func(r *TreeNode, max int) int
	recursion = func(r *TreeNode, max int) (c int) {
		if max < r.Val {
			max = r.Val
		}
		if r.Val == max {
			c++
		}
		if r.Left != nil {
			c += recursion(r.Left, max)
		}
		if r.Right != nil {
			c += recursion(r.Right, max)
		}
		return
	}
	return recursion(root, root.Val)
}

func Test_1544(t *testing.T) {
	NewTestcases(t).
		Add(5, MakeTreeNode("[3,1,4,3,null,1,5,2,5]")).
		Add(4, MakeTreeNode("[3,1,4,3,null,1,5]")).
		Add(3, MakeTreeNode("[3,3,null,4,2]")).
		Add(1, MakeTreeNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := goodNodes(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
