package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2216"
  title: delete-the-middle-node-of-a-linked-list
  lang: golang
  type: large
  inputString: |-
    [1,3,4,7,1,2,6]
    [1,2,3,4]
    [2,1]
    [1]
*/

func deleteMiddle(head *ListNode) *ListNode {
	c := 1
	node := head.Next
	for node != nil {
		node = node.Next
		c++
	}
	if c <= 1 {
		return nil
	}
	mid := (c / 2) - 1
	prev := head
	next := head.Next
	i := 0
	for i < mid && next != nil {
		i++
		prev = next
		next = next.Next
	}
	prev.Next = next.Next
	return head
}

func Test_2216(t *testing.T) {
	NewTestcases(t).
		Add("[1,3,4,1,2,6]", MakeListNode("[1,3,4,7,1,2,6]")).
		Add("[1,2,4]", MakeListNode("[1,2,3,4]")).
		Add("[2]", MakeListNode("[2,1]")).
		Add("[]", MakeListNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := deleteMiddle(td.Input.(*ListNode))

			a.Equal(td.Expectation, MakeListNodeStr(actual))
		})
}
