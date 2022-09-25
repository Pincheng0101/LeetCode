package p2415

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/binarytree"
)

type testCase struct {
	Input  []interface{}
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]interface{}{2, 3, 5, 8, 13, 21, 34}, []int{2, 5, 3, 8, 13, 21, 34}},
	}
}

func TestReverseOddLevels(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := reverseOddLevels(tree.Root)
		assert.Equal(t, testCase.Output, output.GetNodeValuesWithBFS())
	}
}
