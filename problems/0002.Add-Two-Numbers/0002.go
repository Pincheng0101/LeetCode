package p0002

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p1, p2, curr := l1, l2, dummy
	carry := 0
	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
		}
		if p2 != nil {
			v2 = p2.Val
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
