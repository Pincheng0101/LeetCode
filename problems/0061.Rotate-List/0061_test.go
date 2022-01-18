package p0061

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
		{Input{[]int{1, 2, 3, 4, 5}, 2}, []int{4, 5, 1, 2, 3}},
		{Input{[]int{0, 1, 2}, 4}, []int{2, 0, 1}},
		{Input{[]int{1, 2}, 3}, []int{2, 1}},
		{Input{[]int{1}, 1}, []int{1}},
	}
}

func TestRotateRight(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.head)
		output := rotateRight(list.Head.Next, testCase.Input.k)
		outputList := output.Slice()
		assert.Equal(t, testCase.Output, outputList)
	}
}
