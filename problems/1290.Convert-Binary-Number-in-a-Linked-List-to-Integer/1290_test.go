package p1290

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type head []int

type testCase struct {
	Input  head
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 0, 1}, 5},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}, 18880},
		{[]int{0}, 0},
	}
}

func TestGetDecimalValue(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := len(testCase.Input) - 1; i >= 0; i-- {
			list.AddAtHead(testCase.Input[i])
		}
		output := getDecimalValue(list.Head.Next)
		assert.Equal(t, testCase.Output, output)
	}
}
