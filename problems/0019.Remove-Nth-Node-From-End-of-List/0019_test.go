package p0019

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	head []int
	n    int
}

type testCase struct {
	Input  Input
	Output []int // *ListNode (remove nth node)
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 2, 3, 5}},
		{Input{[]int{1}, 1}, []int{}},
		{Input{[]int{1, 2}, 1}, []int{1}},
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.head)
		output := removeNthFromEnd(list.Head.Next, testCase.Input.n)
		outputTestData := output.Slice()

		assert.Equal(t, testCase.Output, outputTestData)
	}
}
