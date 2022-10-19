package leetcode_test

import (
	"math"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "653"
  title: two-sum-iv-input-is-a-bst
  lang: golang
  type: large
  inputString: |-
    [5,3,6,2,4,null,7]
    9
    [5,3,6,2,4,null,7]
    28
*/

func findTarget(root *TreeNode, k int) bool {
	var search func(root *TreeNode, v int) bool
	search = func(root *TreeNode, v int) bool {
		b := root.Val == v
		if !b && root.Left != nil {
			b = b || search(root.Left, v)
		}
		if !b && root.Right != nil {
			b = b || search(root.Right, v)
		}
		return b
	}
	var recursion func(head, curr *TreeNode, k int) bool
	recursion = func(head, curr *TreeNode, k int) bool {
		v := curr.Val
		curr.Val = math.MaxInt
		if search(head, k-v) {
			return true
		}

		if curr.Left != nil {
			if recursion(head, curr.Left, k) {
				return true
			}
		}
		if curr.Right != nil {
			if recursion(head, curr.Right, k) {
				return true
			}
		}
		return false
	}
	return recursion(root, root, k)
}

func Test_653(t *testing.T) {
	type Data struct {
		root *TreeNode
		k    int
	}
	NewTestcases(t).
		Add(true, Data{
			root: MakeTreeNode("[2,1,3]"),
			k:    4,
		}).
		Add(false, Data{
			root: MakeTreeNode("[1]"),
			k:    1,
		}).
		Add(false, Data{
			root: MakeTreeNode("[1]"),
			k:    2,
		}).
		Add(false, Data{
			root: MakeTreeNode("[2,1]"),
			k:    1,
		}).
		Add(false, Data{
			root: MakeTreeNode("[2,1]"),
			k:    2,
		}).
		Add(true, Data{
			root: MakeTreeNode("[2,1]"),
			k:    3,
		}).
		Add(true, Data{
			root: MakeTreeNode("[5,3,6,2,4,null,7]"),
			k:    9,
		}).
		Add(false, Data{
			root: MakeTreeNode("[5,3,6,2,4,null,7]"),
			k:    28,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := findTarget(input.root, input.k)

			a.Equal(td.Expectation, actual)
		})
}
