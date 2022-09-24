package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "113"
  title: path-sum-ii
  lang: golang
  type: large
  inputString: |-
    [5,4,8,11,null,13,4,7,2,null,null,5,1]
    22
    [1,2,3]
    5
    [1,2]
    0
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	var recursion func(*TreeNode, int, int, []int)
	recursion = func(root *TreeNode, targetSum, sum int, v []int) {
		v = append(v, root.Val)
		sum += root.Val
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				res := make([]int, len(v))
				copy(res, v)
				result = append(result, res)
			}
			return
		}
		if root.Left != nil {
			recursion(root.Left, targetSum, sum, v)
		}
		if root.Right != nil {
			recursion(root.Right, targetSum, sum, v)
		}
	}
	recursion(root, targetSum, 0, nil)
	return result
}

func Test_113(t *testing.T) {
	type Data struct {
		root      *TreeNode
		targetSum int
	}
	NewTestcases(t).
		AddExpectation(Make2DMatrixInt("[[5,4,11,2],[5,8,4,5]]")).
		AddInput(Data{
			root:      MakeTreeNode("[5,4,8,11,null,13,4,7,2,null,null,5,1]"),
			targetSum: 22,
		}).
		AddExpectation([][]int{}).
		AddInput(Data{
			root:      MakeTreeNode("[1,2,3]"),
			targetSum: 5,
		}).
		AddExpectation([][]int{}).
		AddInput(Data{
			root:      MakeTreeNode("[1,2]"),
			targetSum: 0,
		}).
		AddExpectation([][]int{}).
		AddInput(Data{
			root:      MakeTreeNode("[]"),
			targetSum: 1,
		}).
		AddExpectation(Make2DMatrixInt("[[-260,-903,-858,-35,817,222,307,-301,-947,-76,445,579,814,-47]]")).
		AddInput(Data{
			root:      MakeTreeNode("[-260,-202,-903,-980,-570,-858,218,764,-300,205,null,-35,null,null,-204,950,-769,258,-652,614,-584,76,817,-192,null,null,-114,880,null,-200,71,671,344,801,193,-18,876,-920,-730,222,679,null,-680,null,null,null,-859,744,-261,692,null,-341,-163,null,null,482,-979,205,null,146,165,801,100,-656,714,-629,995,474,307,-581,-150,-941,null,null,null,-937,-69,-23,82,null,-139,-591,null,-453,-861,-370,null,null,null,216,233,null,430,null,5,-110,null,null,-660,624,-510,-588,null,null,381,null,368,559,null,521,-301,null,522,379,null,null,null,null,456,519,null,null,482,349,null,null,19,null,null,288,-811,null,-372,null,null,-536,null,-404,-457,-740,860,null,null,-636,null,null,342,-874,-462,-504,781,855,-392,null,null,null,406,null,-758,541,null,-947,null,null,null,null,null,-964,null,600,-45,null,null,null,null,null,null,null,null,null,-194,null,null,null,-802,null,null,null,-3,null,-792,672,643,null,14,null,null,489,457,null,null,null,null,412,null,558,null,null,null,null,-846,158,-146,null,null,-76,-650,null,-782,null,-127,null,-678,null,null,null,null,null,null,-464,-426,null,-366,null,null,null,null,null,81,-607,716,null,null,-213,null,379,null,null,null,null,644,445,null,null,-419,-845,-720,null,null,-915,null,null,null,null,null,null,-686,594,-243,null,496,null,907,null,null,null,null,null,null,579,873,702,null,null,null,-834,null,null,null,null,null,-300,-214,-466,null,null,972,null,null,null,814,null,-940,null,763,null,null,null,null,-449,-844,null,null,null,null,-47]"),
			targetSum: -243,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := pathSum(input.root, input.targetSum)

			a.Equal(td.Expectation, actual)
		})
}
