package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "206"
  title: reverse-linked-list
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4,5]
    [1,2]
    []
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	st := arraystack.New()
	curr := head
	for curr != nil {
		st.Push(curr)
		curr = curr.Next
	}
	for !st.Empty() {
		v, _ := st.Pop()
		node := v.(*ListNode)
		if curr == nil {
			head = node
			curr = node
			continue
		}
		curr.Next = node
		curr = node
	}
	curr.Next = nil
	return head
}

func Test_206(t *testing.T) {
	NewTestcases(t).
		Add("[5,4,3,2,1]", MakeListNode("[1,2,3,4,5]")).
		Add("[2,1]", MakeListNode("[1,2]")).
		Add("[]", MakeListNode("[]")).
		Add("[1]", MakeListNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := reverseList(td.Input.(*ListNode))

			a.Equal(td.Expectation, MakeListNodeStr(actual))
		})
}
