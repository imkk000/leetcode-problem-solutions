package leetcode_test

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "775"
  title: n-ary-tree-preorder-traversal
  lang: golang
  type: large
  inputString: |-
    [1,null,3,2,4,null,5,6]
    [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    []
    [1]
    [1,2]
*/

func preorder(root *Node) (v []int) {
	if root == nil {
		return
	}

	st := arraystack.New()
	st.Push(root)
	for !st.Empty() {
		val, _ := st.Pop()
		node := val.(*Node)
		v = append(v, node.Val)
		l := len(node.Children)
		for i := range node.Children {
			st.Push(node.Children[l-i-1])
		}
	}
	return
}

func Test_775(t *testing.T) {
	var intNil []int
	NewTestcases(t).
		AddExpectation(MakeIntSlice("[1,3,5,6,2,4]")).
		AddInput(MakeTreeMultipleNode("[1,null,3,2,4,null,5,6]")).
		AddExpectation(MakeIntSlice("[1,2,3,6,7,11,14,4,8,12,5,9,13,10]")).
		AddInput(MakeTreeMultipleNode("[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]")).
		AddExpectation(intNil).
		AddInput(MakeTreeMultipleNode("[]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := preorder(td.Input.(*Node))

			a.Equal(td.Expectation, actual)
		})
}
