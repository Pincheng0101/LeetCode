package p0082

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
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{[]int{1, 1, 1, 2, 3}, []int{2, 3}},
		{[]int{1, 1}, []int{}},
	}
}

func TestDeleteDuplicates(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input)
		output := deleteDuplicates(list.Head.Next)
		outputAns := output.Slice()
		assert.Equal(t, testCase.Output, outputAns)
	}
}
