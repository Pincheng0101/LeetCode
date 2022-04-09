package p0347

import "container/heap"

type NumFrequent struct {
	v    int
	freq int
}
type MaxFrequentHeap []NumFrequent

func (h MaxFrequentHeap) Len() int           { return len(h) }
func (h MaxFrequentHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxFrequentHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxFrequentHeap) Push(x interface{}) {
	*h = append(*h, x.(NumFrequent))
}

func (h *MaxFrequentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	maxFrequentHeap := &MaxFrequentHeap{}
	heap.Init(maxFrequentHeap)
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for key, cnt := range m {
		heap.Push(maxFrequentHeap, NumFrequent{
			v: key, freq: cnt,
		})
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		topFrequent := heap.Pop(maxFrequentHeap).(NumFrequent)
		ans[i] = topFrequent.v
	}
	return ans
}
