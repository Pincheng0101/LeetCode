package p1046

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, v := range stones {
		heap.Push(maxHeap, v)
	}
	for maxHeap.Len() > 0 {
		if maxHeap.Len() == 1 {
			return heap.Pop(maxHeap).(int)
		}
		first := heap.Pop(maxHeap).(int)
		second := heap.Pop(maxHeap).(int)
		stoneSmashed := first - second
		if stoneSmashed > 0 {
			heap.Push(maxHeap, stoneSmashed)
		}
	}
	return 0
}
