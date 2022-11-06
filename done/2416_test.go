package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2416"
  title: evaluate-boolean-binary-tree
  lang: golang
  type: large
  inputString: |-
    [2,1,3,null,null,0,1]
    [0]
*/

func evaluateTree(root *TreeNode) bool {
	var walk func(ops *arraystack.Stack, root *TreeNode)
	walk = func(ops *arraystack.Stack, root *TreeNode) {
		if root.Val == 2 || root.Val == 3 {
			ops.Push(root)
		}
		if root.Left != nil {
			walk(ops, root.Left)
		}
		if root.Right != nil {
			walk(ops, root.Right)
		}
	}

	ops := arraystack.New()
	walk(ops, root)
	for !ops.Empty() {
		node, _ := ops.Pop()
		v := node.(*TreeNode)
		if v.Val == 2 {
			v.Val = v.Left.Val | v.Right.Val
		}
		if v.Val == 3 {
			v.Val = v.Left.Val & v.Right.Val
		}
	}
	return root.Val == 1
}

func Test_2416(t *testing.T) {
	NewTestcases(t).
		Add(true, MakeTreeNode("[2,1,3,null,null,0,1]")).
		Add(false, MakeTreeNode("[0]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := evaluateTree(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
