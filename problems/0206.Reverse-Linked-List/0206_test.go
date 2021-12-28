package p0206

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	head []int
}
type testCase struct {
	Input  Input
	Output []int
	// Output: *ListNode
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3, 4, 5}}, []int{5, 4, 3, 2, 1}},
		{Input{[]int{1, 2}}, []int{2, 1}},
		{Input{[]int{}}, []int{}},
	}
}

func TestReverseList_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.head); i++ {
			list.AddAtTail(testCase.Input.head[i])
		}
		output := reverseList_1(list.Head.Next)
		outputList := make([]int, 0)
		for output != nil {
			outputList = append(outputList, output.Val)
			output = output.Next
		}
		assert.Equal(t, testCase.Output, outputList)
	}
}

func TestReverseList_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.head); i++ {
			list.AddAtTail(testCase.Input.head[i])
		}
		output := reverseList_2(list.Head.Next)
		outputList := make([]int, 0)
		for output != nil {
			outputList = append(outputList, output.Val)
			output = output.Next
		}
		assert.Equal(t, testCase.Output, outputList)
	}
}
