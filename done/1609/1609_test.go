package leetcode_test

import (
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1609"
  title: find-all-the-lonely-nodes
  lang: golang
  type: large
  inputString: |-
    [1,2,3,null,4]
    [7,1,4,6,null,5,3,null,null,null,null,null,2]
    [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
*/

func getLonelyNodes(root *TreeNode) (s []int) {
	recursion(&s, root)
	return
}

func recursion(s *[]int, root *TreeNode) {
	if root.Left != nil {
		recursion(s, root.Left)
		if root.Right == nil {
			*s = append(*s, root.Left.Val)
		}
	}
	if root.Right != nil {
		recursion(s, root.Right)
		if root.Left == nil {
			*s = append(*s, root.Right.Val)
		}
	}
}

func Test_1609(t *testing.T) {
	var intSliceNil []int
	NewTestcases(t).
		Add(MakeIntSlice("[4]"), MakeTreeNode("[1,2,3,null,4]")).
		Add(MakeIntSlice("[6,2]"), MakeTreeNode("[7,1,4,6,null,5,3,null,null,null,null,null,2]")).
		Add(MakeIntSlice("[77,55,33,66,44,22]"), MakeTreeNode("[[11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]]")).
		Add(MakeIntSlice("[77,55,33,66,44,22]"), MakeTreeNode("[[11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]]")).
		Add(intSliceNil, MakeTreeNode("[1]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := getLonelyNodes(td.Input.(*TreeNode))

			sort.Ints(td.Expectation.([]int))
			sort.Ints(actual)
			a.Equal(td.Expectation, actual)
		})
}
