package p0328

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	oddTail, evenHead := head, head.Next
	curr := evenHead
	for curr != nil && curr.Next != nil {
		oddTail.Next = curr.Next
		oddTail = oddTail.Next
		curr.Next = oddTail.Next
		curr = curr.Next
	}
	oddTail.Next = evenHead
	return head
}
