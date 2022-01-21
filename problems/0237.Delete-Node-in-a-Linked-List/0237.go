package p0237

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	*node = *node.Next
}
