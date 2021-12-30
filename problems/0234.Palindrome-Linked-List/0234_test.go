package p0234

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type testCase struct {
	Input  []int
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1}, true},
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input); i++ {
			list.AddAtTail(testCase.Input[i])
		}
		output := isPalindrome(list.Head.Next)
		assert.Equal(t, testCase.Output, output)
	}
}
