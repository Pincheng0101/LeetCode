package p0203

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	head []int
	val  int
}
type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 6, 3, 4, 5, 6}, 6}, []int{1, 2, 3, 4, 5}},
		{Input{[]int{}, 1}, []int{}},
		{Input{[]int{7, 7, 7, 7}, 7}, []int{}},
	}
}

func TestRemoveElements(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		list := NewSinglyLinkedList()
		list.AddValuesWithSlices(testCase.Input.head)
		output := removeElements(list.Head.Next, testCase.Input.val)
		outputTestData := output.Slice()

		assert.Equal(t, testCase.Output, outputTestData)
	}
}
