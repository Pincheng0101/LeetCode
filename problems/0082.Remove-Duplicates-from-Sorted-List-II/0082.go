package p0082

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head
	duplicate := false
	for cur != nil {
		if cur.Next != nil {
			if cur.Val != cur.Next.Val {
				if duplicate {
					prev.Next = cur.Next
					duplicate = false
				} else {
					prev = prev.Next
				}
			} else {
				duplicate = true
			}
		} else {
			if duplicate {
				prev.Next = nil
			}
		}
		cur = cur.Next
	}

	return dummy.Next
}
