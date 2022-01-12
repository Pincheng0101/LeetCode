package p0142

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	Head []int
	Pos  int
}
type testCase struct {
	Input Input
	// Output: pos index *ListNode
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{3, 2, 0, -4}, 1}},
		{Input{[]int{1, 2}, 0}},
		{Input{[]int{1}, -1}},
	}
}

func TestDetectCycle_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.Head)
		lastNode := list.GetNode(len(testCase.Input.Head) - 1)
		lastNode.Next = list.GetNode(testCase.Input.Pos)
		output := detectCycle_1(list.Head.Next)
		assert.Equal(t, lastNode.Next, output)
	}
}

func TestDetectCycle_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.Head)
		lastNode := list.GetNode(len(testCase.Input.Head) - 1)
		lastNode.Next = list.GetNode(testCase.Input.Pos)
		output := detectCycle_2(list.Head.Next)
		assert.Equal(t, lastNode.Next, output)
	}
}
