package p0148

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
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{[]int{}, []int{}},
	}
}

func TestSortList(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input)
		output := sortList(list.Head.Next)
		outputAns := output.Slice()
		assert.Equal(t, testCase.Output, outputAns)
	}
}
