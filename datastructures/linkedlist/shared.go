package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (this *ListNode) Slice() []int {
	cur := this
	slice := make([]int, 0)
	for cur != nil {
		slice = append(slice, cur.Val)
		cur = cur.Next
	}
	return slice
}
