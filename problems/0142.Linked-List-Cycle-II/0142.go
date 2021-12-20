package p0142

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle_1(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{}, 0)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return cur
		}
		m[cur] = struct{}{}
		cur = cur.Next
	}
	return nil
}

// Fast and Slow Pointer
// slow 和 fast 相遇之後，將 fast 改成 head 直到 head 和 slow 相遇
func detectCycle_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			return nil
		}
		if slow == fast {
			break
		}
	}
	if fast.Next == nil {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
