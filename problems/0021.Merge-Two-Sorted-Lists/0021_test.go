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
		list1.AddValuesWithSlices(testCase.Input.list1)
		list2 := NewSinglyLinkedList()
		list2.AddValuesWithSlices(testCase.Input.list2)
		output := mergeTwoLists_1(list1.Head.Next, list2.Head.Next)
		outputTestData := output.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}

func TestMergeTwoLists_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list1 := NewSinglyLinkedList()
		list1.AddValuesWithSlices(testCase.Input.list1)
		list2 := NewSinglyLinkedList()
		list2.AddValuesWithSlices(testCase.Input.list2)
		output := mergeTwoLists_2(list1.Head.Next, list2.Head.Next)
		outputTestData := output.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}
