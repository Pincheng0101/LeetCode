package p0002

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	l1 []int
	l2 []int
}
type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{2, 4, 3}, []int{5, 6, 4}}, []int{7, 0, 8}},
		{Input{[]int{0}, []int{0}}, []int{0}},
		{Input{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		l1 := NewSinglyLinkedList()
		l1.AddValuesWithSlices(testCase.Input.l1)
		l2 := NewSinglyLinkedList()
		l2.AddValuesWithSlices(testCase.Input.l2)
		output := addTwoNumbers(l1.Head.Next, l2.Head.Next)
		outputTestData := output.Slice()
		assert.Equal(t, testCase.Output, outputTestData)
	}
}
