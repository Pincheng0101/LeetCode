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
func findTilt_1(root *TreeNode) int {
	tile := 0
	sumValues(root, &tile)
	return tile
}

func sumValues(node *TreeNode, tile *int) int {
	if node == nil {
		return 0
	}
	left := sumValues(node.Left, tile)
	right := sumValues(node.Right, tile)
	*tile += abs(left - right)
	return node.Val + left + right
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// https://leetcode-cn.com/problems/binary-tree-tilt/solution/er-cha-shu-de-po-du-by-leetcode-solution-7rha/
func findTilt_2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		ans += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return
}
