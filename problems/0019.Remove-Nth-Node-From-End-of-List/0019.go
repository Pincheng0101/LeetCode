package p0019

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快指針先往前走 n 步，之後快慢指針都各走一步，等到快指針 Next 為 nil 時，慢指針會停在需要刪除的 node 之前
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, head
	for ; n > 1; n-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
