package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "908"
  title: middle-of-the-linked-list
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4,5]
    [1,2,3,4,5,6]
    [1]
    [1,2]
*/

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func Test_908(t *testing.T) {
	NewTestcases(t).
		Add("[3,4,5]", MakeListNode("[1,2,3,4,5]")).
		Add("[4,5,6]", MakeListNode("[1,2,3,4,5,6]")).
		Add("[1]", MakeListNode("[1]")).
		Add("[2]", MakeListNode("[1,2]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := middleNode(td.Input.(*ListNode))

			a.Equal(td.Expectation, MakeListNodeStr(actual))
		})
}

// go:exclude
func middleNodeOld(head *ListNode) *ListNode {
	var b byte
	mid := head
	for head != nil {
		if b++; b&1 == 0 {
			mid = mid.Next
		}
		head = head.Next
	}
	return mid
}
