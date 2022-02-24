package p0148

import (
	"sort"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	list := []int{}
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	sort.Ints(list)

	cur = head
	for _, v := range list {
		cur.Val = v
		cur = cur.Next
	}

	return head
}

// TODO: Merge Sort
