package p0061

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	slow, fast := head, head
	n := 0
	for k > 0 {
		n++
		k--
		fast = fast.Next
		if fast == nil {
			fast = head
			k = k % n
		}
	}
	if fast == slow {
		return head
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next, fast.Next, head = nil, head, slow.Next
	return head
}
