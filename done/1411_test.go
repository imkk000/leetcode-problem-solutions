package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1411"
  title: convert-binary-number-in-a-linked-list-to-integer
  lang: golang
  type: large
  inputString: |-
    [1,0,1]
    [0]
    [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
*/

func getDecimalValue(head *ListNode) int {
	for head.Next != nil {
		head.Next.Val = (head.Val << 1) | head.Next.Val
		head = head.Next
	}
	return head.Val
}

func Test_1411(t *testing.T) {
	NewTestcases(t).
		Add(5, MakeListNode("[1,0,1]")).
		Add(0, MakeListNode("[0]")).
		Add((1<<30)-1, MakeListNode("[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := getDecimalValue(td.Input.(*ListNode))

			a.Equal(td.Expectation, actual)
		})
}
