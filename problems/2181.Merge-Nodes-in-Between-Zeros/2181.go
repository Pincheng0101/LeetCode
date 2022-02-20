package p2181

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	ansCur := dummyHead
	cur := head
	intervalSum := 0
	for cur != nil {
		if cur.Val == 0 {
			if intervalSum != 0 {
				ansCur.Next = &ListNode{Val: intervalSum}
				ansCur = ansCur.Next
			}
			intervalSum = 0
		}
		intervalSum += cur.Val
		cur = cur.Next
	}
	return dummyHead.Next
}
