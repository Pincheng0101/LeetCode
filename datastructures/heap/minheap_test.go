package heap

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	// Heapify the array into a Min Heap
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Add 3，1，2 respectively to the Min Heap
	heap.Push(minHeap, 3)
	heap.Push(minHeap, 1)
	heap.Push(minHeap, 2)

	// Check all elements in the Min Heap, the result is [1, 3, 2]
	t.Log("minHeap: ", minHeap)

	// Get the top element of the Min Heap
	peekNum := (*minHeap)[0]

	// The result is 1
	t.Log("peek number: ", peekNum)

	// Delete the top element in the Min Heap
	popNum := heap.Pop(minHeap)

	// The result is 1
	t.Log("pop number: ", popNum)

	// Check the top element after deleting 1, the result is 2
	t.Log("peek number: ", (*minHeap)[0])

	// Check all elements in the Min Heap, the result is [2,3]
	t.Log("minHeap: ", minHeap)

	// Check the number of elements in the Min Heap
	// Which is also the length of the Min Heap
	size := minHeap.Len()

	// The result is 2
	t.Log("minHeap size: ", size)
}
