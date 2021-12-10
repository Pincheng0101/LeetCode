package p0563

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(sumNode(root.Left)-sumNode(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}

func sumNode(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + sumNode(node.Left) + sumNode(node.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
