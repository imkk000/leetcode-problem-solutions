package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) {
	q := NewQueue[*TreeNode]()
	q.Enqueue(root)
	for !q.Empty() {
		node := q.Dequeue()
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
}

func TestBFS(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	bfs(root)
}

func TestQueueTreeNode(t *testing.T) {
	expectation := []*TreeNode{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
	}

	q := NewQueue[*TreeNode]()
	for _, n := range expectation {
		q.Enqueue(n)
	}
	actual := make([]*TreeNode, 0)
	for !q.Empty() {
		actual = append(actual, q.Dequeue())
	}

	assert.Equal(t, expectation, actual)
}

func TestQueueInt(t *testing.T) {
	expectation := []int{1, 2, 3, 4, 5}

	q := NewQueue[int]()
	for _, n := range expectation {
		q.Enqueue(n)
	}
	actual := make([]int, 0)
	for !q.Empty() {
		actual = append(actual, q.Dequeue())
	}

	assert.Equal(t, expectation, actual)
}

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.q = append(q.q, n)
}

func (q *Queue[T]) Dequeue() (v T) {
	v = q.q[0]
	q.q = q.q[1:]
	return
}

func (q *Queue[T]) Empty() bool {
	return len(q.q) == 0
}
