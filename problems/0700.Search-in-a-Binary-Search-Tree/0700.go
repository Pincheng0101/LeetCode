package p0700

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	l := searchBST(root.Left, val)
	if l != nil {
		return l
	}
	r := searchBST(root.Right, val)
	if r != nil {
		return r
	}
	return nil
}
