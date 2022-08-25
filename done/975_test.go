package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "975"
  title: range-sum-of-bst
  lang: golang
  type: large
  inputString: |-
    [10,5,15,3,7,null,18]
    7
    15
    [10,5,15,3,7,13,18,1,null,6]
    6
    10
    [18,9,27,6,15,24,30,3,null,12,null,21]
    18
    24
*/

func rangeSumBST(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if high < root.Val {
		return rangeSumBST(root.Left, low, high)
	}
	if low > root.Val {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val +
		rangeSumBST(root.Left, low, high) +
		rangeSumBST(root.Right, low, high)
}

func Test_975(t *testing.T) {
	type Data struct {
		root      *TreeNode
		low, high int
	}
	NewTestcases(t).
		AddExpectation(32).
		AddInput(Data{
			root: MakeTreeNode("[10,5,15,3,7,null,18]"),
			low:  7,
			high: 15,
		}).
		AddExpectation(23).
		AddInput(Data{
			root: MakeTreeNode("[10,5,15,3,7,13,18,1,null,6]"),
			low:  6,
			high: 10,
		}).
		AddExpectation(0).
		AddInput(Data{
			root: MakeTreeNode("[1,2,3,4,5,6,7,8,9]"),
			low:  99,
			high: 100,
		}).
		AddExpectation(78).
		AddInput(Data{
			root: MakeTreeNode("[10,5,15,3,7,13,18,1,null,6]"),
			low:  1,
			high: 18,
		}).
		AddExpectation(63).
		AddInput(Data{
			root: MakeTreeNode("[18,9,27,6,15,24,30,3,null,12,null,21]"),
			low:  18,
			high: 24,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := rangeSumBST(input.root, input.low, input.high)

			a.Equal(td.Expectation, actual)
		})
}
