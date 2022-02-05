package p0023

import (
	"container/heap"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, listNode := range lists {
		if listNode != nil {
			heap.Push(pq, listNode)
		}
	}

	var dummyHead *ListNode = &ListNode{}
	var tail *ListNode = dummyHead
	for pq.Len() != 0 {
		minNode := heap.Pop(pq).(*ListNode)
		tail.Next = minNode
		tail = tail.Next
		if minNode.Next != nil {
			heap.Push(pq, minNode.Next)
		}
	}

	return dummyHead.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
