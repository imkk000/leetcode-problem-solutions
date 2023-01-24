package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "783"
  title: search-in-a-binary-search-tree
  lang: golang
  type: large
  inputString: |-
    [4,2,7,1,3]
    2
    [4,2,7,1,3]
    5
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	st := arraystack.New()
	st.Push(root)
	for !st.Empty() {
		v, _ := st.Pop()
		node := v.(*TreeNode)
		if node.Val == val {
			return node
		}
		if node.Left != nil {
			st.Push(node.Left)
		}
		if node.Right != nil {
			st.Push(node.Right)
		}
	}
	return nil
}

func Test_783(t *testing.T) {
	type Data struct {
		root *TreeNode
		val  int
	}
	NewTestcases(t).
		Add(MakeTreeNode("[2,1,3]"), Data{
			root: MakeTreeNode("[4,2,7,1,3]"),
			val:  2,
		}).
		Add(TreeNodeNil, Data{
			root: MakeTreeNode("[4,2,7,1,3]"),
			val:  5,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			root, val := input.root, input.val
			actual := searchBST(root, val)

			a.Equal(td.Expectation, actual)
		})
}
