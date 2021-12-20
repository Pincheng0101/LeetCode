package p0141

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle_1(head *ListNode) bool {
	m := make(map[*ListNode]struct{}, 0)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}

// 一個一次走兩步，一個一次走一步，如果是 cycle 遲早會追到
func hasCycle_2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
	}
	return false

}
