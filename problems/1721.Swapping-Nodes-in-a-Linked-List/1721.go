package p1721

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	node1, node2, node2fast := head, head, head
	for i := 1; i < k; i++ {
		node1 = node1.Next
		node2fast = node2fast.Next
	}
	node2fast = node2fast.Next

	for node2fast != nil {
		node2 = node2.Next
		node2fast = node2fast.Next
	}

	node1.Val, node2.Val = node2.Val, node1.Val

	return head
}
