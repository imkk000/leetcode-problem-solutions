package leetcode_test

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "606"
  title: construct-string-from-binary-tree
  lang: golang
  type: large
  inputString: |-
    [1,2,3,4]
    [1,2,3,null,4]
    [2]
    [1,2,3]
    [1,null,2]
*/

func tree2str(root *TreeNode) string {
	s := strconv.Itoa(root.Val)
	if root.Left != nil {
		s += fmt.Sprintf("(%s)", tree2str(root.Left))
	}
	if root.Right != nil {
		if root.Left == nil {
			s += "()"
		}
		s += fmt.Sprintf("(%s)", tree2str(root.Right))
	}
	return s
}

func Test_606(t *testing.T) {
	NewTestcases(t).
		Add("2", MakeTreeNode("[2]")).
		Add("1(2)(3)", MakeTreeNode("[1,2,3]")).
		Add("1()(2)", MakeTreeNode("[1,null,2]")).
		Add("1(2(4))(3)", MakeTreeNode("[1,2,3,4]")).
		Add("1(2()(4))(3)", MakeTreeNode("[1,2,3,null,4]")).
		Add("-1000", MakeTreeNode("[-1000]")).
		Add("1000", MakeTreeNode("[1000]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := tree2str(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}
