package p2415

import . "github.com/pincheng0101/leetcode/datastructures/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverseOddLevels(root *TreeNode) *TreeNode {
	var level int = 0
	var nodes []*TreeNode = make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) != 0 {
		if level%2 == 1 {
			n := len(nodes)
			for i := 0; i < n/2; i++ {
				nodes[i].Val, nodes[n-i-1].Val = nodes[n-i-1].Val, nodes[i].Val
			}
		}

		var nextNodes []*TreeNode = make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		nodes = nextNodes

		level++
	}

	return root
}
