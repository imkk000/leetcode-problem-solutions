package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "100"
  title: same-tree
  lang: golang
  type: large
  inputString: |-
    [1,2,3]
    [1,2,3]
    [1,2]
    [1,null,2]
    [1,2,1]
    [1,1,2]
    []
    []
    [1]
    [1]
    [2]
    [1]
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	st := NewStack()
	st.Push([]*TreeNode{p, q})
	for !st.Empty() {
		node := st.Pop()
		if node[0] == nil && node[1] == nil {
			continue
		}
		if node[0] != nil && node[1] == nil ||
			node[0] == nil && node[1] != nil ||
			node[0].Val != node[1].Val {
			return false
		}
		st.Push([]*TreeNode{node[0].Left, node[1].Left})
		st.Push([]*TreeNode{node[0].Right, node[1].Right})
	}
	return true
}

type Stack [][]*TreeNode

func NewStack() Stack {
	return make([][]*TreeNode, 0)
}

func (s *Stack) Push(v []*TreeNode) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (v []*TreeNode) {
	if s.Empty() {
		return nil
	}
	l := len(*s) - 1
	v = (*s)[l]
	(*s)[l] = nil
	*s = (*s)[:l]
	return
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// go:exclude
const (
	round = 1e2
)

func BenchmarkImpl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := NewStack()
		for i := 0; i < round; i++ {
			st.Push([]*TreeNode{nil, nil})
			_ = st.Pop()
		}
	}
}

func BenchmarkGoDS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := arraystack.New()
		for i := 0; i < round; i++ {
			st.Push([]*TreeNode{nil, nil})
			_, _ = st.Pop()
		}
	}
}

func Test_100(t *testing.T) {
	type Data struct {
		p, q *TreeNode
	}
	NewTestcases(t).
		Add(true, Data{
			p: MakeTreeNode("[1,2,3]"),
			q: MakeTreeNode("[1,2,3]"),
		}).
		Add(false, Data{
			p: MakeTreeNode("[1,2]"),
			q: MakeTreeNode("[1,null,2]"),
		}).
		Add(false, Data{
			p: MakeTreeNode("[1,2,1]"),
			q: MakeTreeNode("[1,1,2]"),
		}).
		Add(true, Data{
			p: MakeTreeNode("[]"),
			q: MakeTreeNode("[]"),
		}).
		Add(true, Data{
			p: MakeTreeNode("[1]"),
			q: MakeTreeNode("[1]"),
		}).
		Add(false, Data{
			p: MakeTreeNode("[2]"),
			q: MakeTreeNode("[1]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			p, q := input.p, input.q
			actual := isSameTree(p, q)

			a.Equal(td.Expectation, actual)
		})
}
