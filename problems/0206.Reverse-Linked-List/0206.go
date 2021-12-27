package p0206

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	newHead := head
	if head != nil {
		for head.Next != nil {
			temp := head.Next
			head.Next = head.Next.Next
			temp.Next = newHead
			newHead = temp
		}
	}
	return newHead
}
