package p2181

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type testCase struct {
	Input  []int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]int{0, 3, 1, 0, 4, 5, 2, 0}, []int{4, 11}},
		{[]int{0, 1, 0, 3, 0, 2, 2, 0}, []int{1, 3, 4}},
	}
}

func TestMergeNodes(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input)
		output := mergeNodes(list.Head)
		outputAns := output.Slice()
		assert.Equal(t, testCase.Output, outputAns)
	}
}
