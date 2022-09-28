package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "19"
  title: remove-nth-node-from-end-of-list
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4,5]
    2
    [1]
    1
    [1,2]
    1
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	for node != nil {
		n--
		node = node.Next
	}

	var prev *ListNode
	node = head
	for n < 0 && node.Next != nil {
		n++
		prev = node
		node = node.Next
	}

	if prev == nil {
		head = head.Next
	} else {
		prev.Next = node.Next
	}
	return head
}

func Test_19(t *testing.T) {
	type Data struct {
		head *ListNode
		n    int
	}
	NewTestcases(t).
		AddExpectation("[1,2,3,5]").
		AddInput(Data{
			head: MakeListNode("[1,2,3,4,5]"),
			n:    2,
		}).
		AddExpectation("[]").
		AddInput(Data{
			head: MakeListNode("[1]"),
			n:    1,
		}).
		AddExpectation("[1]").
		AddInput(Data{
			head: MakeListNode("[1,2]"),
			n:    1,
		}).
		AddExpectation("[2]").
		AddInput(Data{
			head: MakeListNode("[1,2]"),
			n:    2,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := removeNthFromEnd(input.head, input.n)

			a.Equal(td.Expectation, MakeListNodeStr(actual))
		})
}
