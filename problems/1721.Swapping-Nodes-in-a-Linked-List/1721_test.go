package p1721

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	head []int
	k    int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 4, 3, 2, 5}},
		{Input{[]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}, 5}, []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}},
	}
}

func TestSwapNodes(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.head)
		output := swapNodes(list.Head.Next, testCase.Input.k)
		outputList := output.Slice()
		assert.Equal(t, testCase.Output, outputList)
	}
}
