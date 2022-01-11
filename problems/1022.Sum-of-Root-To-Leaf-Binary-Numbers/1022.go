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
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var sum int = 0
	if root.Left != nil {
		root.Left.Val += root.Val << 1
		sum += sumRootToLeaf(root.Left)
	}
	if root.Right != nil {
		root.Right.Val += root.Val << 1
		sum += sumRootToLeaf(root.Right)
	}
	return sum
}
