package p0382

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

func TestGetRandom(t *testing.T) {
	list := NewSinglyLinkedList()
	list.AddValuesWithSlices([]int{1, 2, 3})
	solution := Constructor(list.Head)

	// TODO: Unit Testing Randomness
	assert.Equal(t, solution.GetRandom(), 1)
	assert.Equal(t, solution.GetRandom(), 2)
	assert.Equal(t, solution.GetRandom(), 3)
}
