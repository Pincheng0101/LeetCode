package p0141

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
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{3, 2, 0, -4}, 1}, true},
		{Input{[]int{1, 2}, 0}, true},
		{Input{[]int{1}, -1}, false},
	}
}

func TestHasCycle_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.Head); i++ {
			list.AddAtTail(testCase.Input.Head[i])
		}
		lastNode := list.GetNode(len(testCase.Input.Head) - 1)
		lastNode.Next = list.GetNode(testCase.Input.Pos)
		output := hasCycle_1(list.Head)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestHasCycle_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.Head); i++ {
			list.AddAtTail(testCase.Input.Head[i])
		}
		lastNode := list.GetNode(len(testCase.Input.Head) - 1)
		lastNode.Next = list.GetNode(testCase.Input.Pos)
		output := hasCycle_2(list.Head)
		assert.Equal(t, testCase.Output, output)
	}
}
