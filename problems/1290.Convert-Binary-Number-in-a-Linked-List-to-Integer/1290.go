package p1290

import . "github.com/pincheng0101/leetcode/datastructures/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	cur := head
	num := 0
	for cur != nil {
		num = num*2 + cur.Val
		cur = cur.Next
	}
	return num
}
