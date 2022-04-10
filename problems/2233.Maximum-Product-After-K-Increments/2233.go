package p2233

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, v := range nums {
		heap.Push(minHeap, v)
	}
	for k > 0 {
		v := heap.Pop(minHeap).(int)
		heap.Push(minHeap, v+1)
		k--
	}
	ans := 1
	for _, v := range *minHeap {
		ans = (ans * v) % 1000000007
	}
	return ans
}
