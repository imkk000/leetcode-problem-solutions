package leetcode_test

import (
	_ "embed"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1568"
  title: pseudo-palindromic-paths-in-a-binary-tree
  lang: golang
  type: large
  inputString: |-
    [2,3,1,3,1,null,1]
    [2,1,1,1,3,null,null,null,null,null,1]
    [9]
*/

func pseudoPalindromicPaths(root *TreeNode) (c int) {
	var recursion func(root *TreeNode, s [10]bool) int
	recursion = func(root *TreeNode, s [10]bool) (c int) {
		s[root.Val-1] = !s[root.Val-1]
		if root.Left == nil && root.Right == nil {
			var v byte
			for _, b := range s {
				if b {
					v++
				}
				if v > 1 {
					return
				}
			}
			return 1
		}
		if root.Left != nil {
			c += recursion(root.Left, s)
		}
		if root.Right != nil {
			c += recursion(root.Right, s)
		}
		return
	}
	return recursion(root, [10]bool{})
}

func Test_1568(t *testing.T) {
	NewTestcases(t).
		Add(2, MakeTreeNode("[2,3,1,3,1,null,1]")).
		Add(1, MakeTreeNode("[2,1,1,1,3,null,null,null,null,null,1]")).
		Add(1, MakeTreeNode("[9]")).
		Add(885, MakeTreeNode(input1568)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := pseudoPalindromicPaths(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}

// go:exclude
//
//go:embed testdata/1568_1_in
var input1568 string
