package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "234"
  title: palindrome-linked-list
  lang: golang
  type: large
  inputString: |-
    [1,2,2,1]
    [1,2]
*/

func isPalindrome(head *ListNode) bool {
	var v []int
	for head != nil {
		v = append(v, head.Val)
		head = head.Next
	}
	n := len(v)
	for i := 0; i <= n/2; i++ {
		if v[i] != v[n-i-1] {
			return false
		}
	}
	return true
}

func Test_234(t *testing.T) {
	NewTestcases(t).
		Add(false, "[1,1,2,1]").
		Add(true, "[1,2,1]").
		Add(true, "[1,2,2,1]").
		Add(false, "[1,2]").
		Add(true, "[1]").
		Each(func(a *assert.Assertions, td TestData) {
			input := MakeListNode(td.Input.(string))
			actual := isPalindrome(input)

			a.Equal(td.Expectation, actual)
		})
}
