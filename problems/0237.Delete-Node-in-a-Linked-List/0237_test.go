package p0237

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	head []int
	node int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{4, 5, 1, 9}, 5}, []int{4, 1, 9}},
		{Input{[]int{4, 5, 1, 9}, 1}, []int{4, 5, 9}},
	}
}

func TestDeleteNode(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.head)
		node := list.GetNodeByValue(testCase.Input.node)
		deleteNode(node)
		outputTestData := list.Head.Next.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}
