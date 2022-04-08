package p0703

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

func (h *MinHeap) Peek() interface{} {
	old := *h
	return old[0]
}

type KthLargest struct {
	minHeap *MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	kthLargest := KthLargest{minHeap: minHeap, k: k}
	for _, v := range nums {
		kthLargest.Add(v)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return this.minHeap.Peek().(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
