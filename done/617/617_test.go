package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "617"
  title: merge-two-binary-trees
  lang: golang
  type: large
  inputString: |-
    [1,3,2,5]
    [2,1,3,null,4,null,7]
    [1]
    [1,2]
    [1]
    [2]
    [1,2]
    [1,3]
    [1]
    []
    []
    [1]
    []
    []
    [1,2,null,3]
    [1,2,null,null,4]
    [1,2,null,3]
    [1,null,2,null,3]
*/

type MergeTreeNodes struct {
	node1, node2    *TreeNode
	prevNode        *TreeNode
	isLeft, isRight bool
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 != nil {
		return root2
	}

	st := NewStack()
	st.Push(MergeTreeNodes{
		node1: root1,
		node2: root2,
	})
	for !st.Empty() {
		v, _ := st.Pop()
		mtNodes := v.(MergeTreeNodes)
		if mtNodes.node1 == nil && mtNodes.node2 == nil {
			continue
		}
		if mtNodes.node1 != nil && mtNodes.node2 != nil {
			mtNodes.node1.Val += mtNodes.node2.Val

			st.Push(MergeTreeNodes{
				node1:    mtNodes.node1.Left,
				node2:    mtNodes.node2.Left,
				prevNode: mtNodes.node1,
				isLeft:   true,
				isRight:  false,
			})
			st.Push(MergeTreeNodes{
				node1:    mtNodes.node1.Right,
				node2:    mtNodes.node2.Right,
				prevNode: mtNodes.node1,
				isLeft:   false,
				isRight:  true,
			})
		}
		if mtNodes.node1 == nil && mtNodes.node2 != nil {
			if mtNodes.isLeft {
				mtNodes.prevNode.Left = mtNodes.node2
			}
			if mtNodes.isRight {
				mtNodes.prevNode.Right = mtNodes.node2
			}
		}
	}
	return root1
}

type Stack struct {
	v []MergeTreeNodes
	l int
}

func NewStack() Stack {
	return Stack{
		v: make([]MergeTreeNodes, 0),
	}
}

func (s *Stack) Push(m MergeTreeNodes) {
	s.v = append(s.v, m)
	s.l++
}

func (s *Stack) Pop() (m interface{}, d struct{}) {
	s.l--
	m = s.v[s.l]
	s.v = s.v[:s.l]
	return
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return s.l
}

func Test_617(t *testing.T) {
	type Data struct {
		root1, root2 *TreeNode
	}
	NewTestcases(t).
		Add(MakeTreeNode("[3,4,5,5,4,null,7]"), Data{
			root1: MakeTreeNode("[1,3,2,5]"),
			root2: MakeTreeNode("[2,1,3,null,4,null,7]"),
		}).
		Add(MakeTreeNode("[2,2]"), Data{
			root1: MakeTreeNode("[1]"),
			root2: MakeTreeNode("[1,2]"),
		}).
		Add(MakeTreeNode("[3]"), Data{
			root1: MakeTreeNode("[1]"),
			root2: MakeTreeNode("[2]"),
		}).
		Add(MakeTreeNode("[2,2,3]"), Data{
			root1: MakeTreeNode("[1,2]"),
			root2: MakeTreeNode("[1,null,3]"),
		}).
		Add(MakeTreeNode("[1]"), Data{
			root1: MakeTreeNode("[1]"),
			root2: MakeTreeNode("[]"),
		}).
		Add(MakeTreeNode("[1]"), Data{
			root1: MakeTreeNode("[]"),
			root2: MakeTreeNode("[1]"),
		}).
		Add(MakeTreeNode("[]"), Data{
			root1: MakeTreeNode("[]"),
			root2: MakeTreeNode("[]"),
		}).
		Add(MakeTreeNode("[2,4,null,3,4]"), Data{
			root1: MakeTreeNode("[1,2,null,3]"),
			root2: MakeTreeNode("[1,2,null,null,4]"),
		}).
		Add(MakeTreeNode("[2,2,2,3,null,null,3]"), Data{
			root1: MakeTreeNode("[1,2,null,3]"),
			root2: MakeTreeNode("[1,null,2,null,3]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			root1, root2 := input.root1, input.root2
			actual := mergeTrees(root1, root2)

			a.Equal(td.Expectation, actual)
		})
}
