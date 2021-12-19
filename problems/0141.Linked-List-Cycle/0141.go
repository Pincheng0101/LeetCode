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
	if head == nil {
		return false
	}
	first := head
	second := head.Next
	for second != nil {
		if first == second {
			return true
		}
		if second.Next == nil {
			return false
		}
		first = first.Next
		second = second.Next.Next
	}
	return false

}
