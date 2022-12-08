package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "904"
  title: leaf-similar-trees
  lang: golang
  type: large
  inputString: |-
    [3,5,1,6,2,9,8,null,null,7,4]
    [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
    [1,2,3]
    [1,3,2]
    [1,3]
    [1,null,3]
    [4,2,6,1,3,5,7]
    [4,2,6,null,3,5,7]
*/

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	st := arraystack.New()
	st.Push(root1)
	var s []int
	for !st.Empty() {
		v, _ := st.Pop()
		node := v.(*TreeNode)
		if node.Left == nil && node.Right == nil {
			s = append(s, node.Val)
			continue
		}
		if node.Left != nil {
			st.Push(node.Left)
		}
		if node.Right != nil {
			st.Push(node.Right)
		}
	}
	st.Push(root2)
	l := len(s)
	var i int
	for !st.Empty() {
		v, _ := st.Pop()
		node := v.(*TreeNode)
		if node.Left == nil && node.Right == nil {
			if i >= l {
				return false
			}
			if s[i] != node.Val {
				return false
			}
			i++
			continue
		}
		if node.Left != nil {
			st.Push(node.Left)
		}
		if node.Right != nil {
			st.Push(node.Right)
		}
	}
	return i == l
}

func Test_904(t *testing.T) {
	type Data struct {
		root1, root2 *TreeNode
	}
	NewTestcases(t).
		Add(true, Data{
			root1: MakeTreeNode("[3,5,1,6,2,9,8,null,null,7,4]"),
			root2: MakeTreeNode("[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]"),
		}).
		Add(false, Data{
			root1: MakeTreeNode("[1,2,3]"),
			root2: MakeTreeNode("[1,3,2]"),
		}).
		Add(false, Data{
			root1: MakeTreeNode("[1,2]"),
			root2: MakeTreeNode("[1,2,3]"),
		}).
		Add(true, Data{
			root1: MakeTreeNode("[1]"),
			root2: MakeTreeNode("[1]"),
		}).
		Add(false, Data{
			root1: MakeTreeNode("[2]"),
			root2: MakeTreeNode("[1]"),
		}).
		Add(true, Data{
			root1: MakeTreeNode("[1,3]"),
			root2: MakeTreeNode("[1,null,3]"),
		}).
		Add(false, Data{
			root1: MakeTreeNode("[4,2,6,1,3,5,7]"),
			root2: MakeTreeNode("[4,2,6,null,3,5,7]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			root1, root2 := input.root1, input.root2
			actual := leafSimilar(root1, root2)

			a.Equal(td.Expectation, actual)
		})
}
