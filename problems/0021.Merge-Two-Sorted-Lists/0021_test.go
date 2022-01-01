package p0021

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	list1 []int
	list2 []int
}
type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 4}, []int{1, 3, 4}}, []int{1, 1, 2, 3, 4, 4}},
		{Input{[]int{}, []int{}}, []int{}},
		{Input{[]int{}, []int{0}}, []int{0}},
	}
}

func TestMergeTwoLists_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list1 := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.list1); i++ {
			list1.AddAtTail(testCase.Input.list1[i])
		}
		list2 := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.list2); i++ {
			list2.AddAtTail(testCase.Input.list2[i])
		}
		output := mergeTwoLists_1(list1.Head.Next, list2.Head.Next)
		outputTestData := output.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}

func TestMergeTwoLists_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list1 := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.list1); i++ {
			list1.AddAtTail(testCase.Input.list1[i])
		}
		list2 := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.list2); i++ {
			list2.AddAtTail(testCase.Input.list2[i])
		}
		output := mergeTwoLists_2(list1.Head.Next, list2.Head.Next)
		outputTestData := output.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}
