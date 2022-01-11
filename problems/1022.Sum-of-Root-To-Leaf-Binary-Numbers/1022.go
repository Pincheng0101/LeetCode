package p1022

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Easy Solution，但是會修改到原本的數值
func sumRootToLeaf_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var sum int = 0
	if root.Left != nil {
		root.Left.Val += root.Val << 1
		sum += sumRootToLeaf_1(root.Left)
	}
	if root.Right != nil {
		root.Right.Val += root.Val << 1
		sum += sumRootToLeaf_1(root.Right)
	}
	return sum
}

// 多一個值來存 currSum，不影響原本數值
func sumRootToLeaf_2(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, currSum int) int {
	if root == nil {
		return 0
	}

	currSum = (currSum << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		return currSum
	}

	return dfs(root.Left, currSum) + dfs(root.Right, currSum)
}
