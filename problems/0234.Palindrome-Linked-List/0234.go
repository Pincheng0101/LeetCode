package p0234

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	newHead := head
	slow, fast := head, head
	// 反轉前半段
	// ex. 1 -> 2 -> 2 -> 1 to 2 -> 1 -> 2 -> 1
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow, slow.Next, newHead = slow.Next, newHead, slow
	}
	// 奇數節點需要做處理，跳過正中間的節點
	if fast != nil {
		slow = slow.Next
	}
	// 比對反轉後的前半段和後半段節點數值
	for slow != nil {
		if newHead.Val != slow.Val {
			return false
		}
		newHead = newHead.Next
		slow = slow.Next
	}
	return true
}
