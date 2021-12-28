package p0206

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	for head.Next != nil {
		curr := head.Next
		head.Next = head.Next.Next
		curr.Next = newHead
		newHead = curr
	}
	return newHead
}

// 利用 golang 的特性，在一行中直接交換位置，確保結點位置不會遺失
func reverseList_2(head *ListNode) *ListNode {
	newHead := (*ListNode)(nil)
	for curr := head; curr != nil; {
		curr, curr.Next, newHead = curr.Next, newHead, curr
	}
	return newHead
}
