package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "237"
  title: delete-node-in-a-linked-list
  lang: golang
  type: large
  inputString: |-
    [4,5,1,9]
    5
    [4,5,1,9]
    1
    [1,2]
    1
*/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func Test_237(t *testing.T) {
	type Data struct {
		head *ListNode
		node int
	}
	NewTestcases(t).
		Add("[4,1,9]", Data{
			head: MakeListNode("[4,5,1,9]"),
			node: 5,
		}).
		Add("[4,5,9]", Data{
			head: MakeListNode("[4,5,1,9]"),
			node: 1,
		}).
		Add("[2]", Data{
			head: MakeListNode("[1,2]"),
			node: 1,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			head, node := input.head, input.node
			next := head
			for next.Val != node {
				next = next.Next
			}
			deleteNode(next)

			a.Equal(td.Expectation, MakeListNodeStr(head))
		})
}
