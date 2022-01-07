package p0382

import (
	"math/rand"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	data []int
}

func Constructor(head *ListNode) Solution {
	data := make([]int, 0)
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	return Solution{data: data}
}

func (this *Solution) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
