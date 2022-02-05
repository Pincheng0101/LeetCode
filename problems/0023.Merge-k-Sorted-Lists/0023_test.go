package p0023

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type testCase struct {
	Input  [][]int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
		{[][]int{{}, {1}}, []int{1}},
	}
}

func TestFunction(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		lists := []*ListNode{}
		for _, slice := range testCase.Input {
			list := NewSinglyLinkedList()
			list.AddValuesWithSlices(slice)
			lists = append(lists, list.Head.Next)
		}
		output := mergeKLists(lists)
		ans := output.Slice()
		assert.Equal(t, testCase.Output, ans)
	}
}
