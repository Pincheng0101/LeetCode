package p0104

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth
		}
		depth++
		leftMaxDepth := dfs(node.Left, depth)
		rightMaxDepth := dfs(node.Right, depth)
		return max(leftMaxDepth, rightMaxDepth)
	}
	return dfs(root, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
