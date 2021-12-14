package p0938

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 要注意 binary tree 的特性，比 node 大的值放右邊，小的放左邊
func rangeSumBST(root *TreeNode, low int, high int) (ans int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val >= low && node.Val <= high {
			ans += node.Val
		}
		if node.Val >= low {
			dfs(node.Left)
		}
		if node.Val <= high {
			dfs(node.Right)
		}
		return
	}
	dfs(root)
	return
}
