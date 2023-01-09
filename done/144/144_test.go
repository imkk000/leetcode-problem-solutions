package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "144"
  title: binary-tree-preorder-traversal
  lang: golang
  type: large
  inputString: |-
    [1,null,2,3]
    []
    [1]
*/

func preorderTraversal(root *TreeNode) (s []int) {
	if root == nil {
		return
	}

	st := make(Stack, 0)
	s = append(s, root.Val)
	walk := func(root *TreeNode) {
		st.Push(root)
		for !st.Empty() {
			node := st.Pop()
			s = append(s, node.Val)
			if node.Right != nil {
				st.Push(node.Right)
			}
			if node.Left != nil {
				st.Push(node.Left)
			}
		}
	}
	if root.Left != nil {
		walk(root.Left)
	}
	if root.Right != nil {
		walk(root.Right)
	}
	return
}

type Stack []*TreeNode

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(v *TreeNode) {
	*s = append([]*TreeNode{v}, (*s)...)
}

func (s *Stack) Pop() (node *TreeNode) {
	if s.Empty() {
		return
	}
	node = (*s)[0]
	(*s)[0] = nil
	*s = (*s)[1:]
	return
}

func (s *Stack) Empty() bool {
	return len((*s)) == 0
}

func Test_144(t *testing.T) {
	var intSliceNil []int
	NewTestcases(t).
		Add(intSliceNil, MakeTreeNode("[]")).
		Add(MakeIntSlice("[1]"), MakeTreeNode("[1]")).
		Add(MakeIntSlice("[1,2,3]"), MakeTreeNode("[1,null,2,3]")).
		Add(MakeIntSlice("[1,2,3]"), MakeTreeNode("[1,2,3]")).
		Add(MakeIntSlice("[1,2,3,4,5]"), MakeTreeNode("[1,2,3,null,null,4,5]")).
		Add(MakeIntSlice("[1,2,4,3,5]"), MakeTreeNode("[1,2,3,null,4,5]")).
		Add(MakeIntSlice("[1,2,4,8,5,3,6,9,7,10]"), MakeTreeNode("[1,2,3,4,5,6,7,8,null,null,null,9,null,10]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := preorderTraversal(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
