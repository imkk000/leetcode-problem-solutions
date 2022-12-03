package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "21"
  title: merge-two-sorted-lists
  lang: golang
  type: large
  inputString: |-
    [1,2,4]
    [1,3,4]
    []
    []
    []
    [0]
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	var newList, currNode *ListNode
	for l1 != nil {
		node := &ListNode{
			Val: l1.Val,
		}
		if newList == nil {
			newList = node
			currNode = newList
		} else {
			currNode.Next = node
			currNode = currNode.Next
		}
		mx := 101
		if l1.Next != nil {
			mx = l1.Next.Val
		}
		for l2 != nil {
			if l2.Val > mx {
				break
			}
			node = &ListNode{
				Val: l2.Val,
			}
			currNode.Next = node
			currNode = currNode.Next
			l2 = l2.Next
		}
		l1 = l1.Next
	}
	return newList
}

func Test_21(t *testing.T) {
	type Data struct {
		l1, l2 *ListNode
	}
	NewTestcases(t).
		Add("[1,1,2,3,4,4]", Data{
			l1: MakeListNode("[1,2,4]"),
			l2: MakeListNode("[1,3,4]"),
		}).
		Add("[]", Data{
			l1: MakeListNode("[]"),
			l2: MakeListNode("[]"),
		}).
		Add("[0]", Data{
			l1: MakeListNode("[]"),
			l2: MakeListNode("[0]"),
		}).
		Add("[1]", Data{
			l1: MakeListNode("[1]"),
			l2: MakeListNode("[]"),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			l1, l2 := input.l1, input.l2
			actual := mergeTwoLists(l1, l2)

			a.Equal(td.Expectation, MakeListNodeStr(actual))
		})
}
