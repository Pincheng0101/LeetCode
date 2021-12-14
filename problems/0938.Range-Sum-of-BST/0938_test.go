package p0938

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/binarytree"
)

type Input struct {
	root []interface{}
	low  int
	high int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{
			Input: Input{
				root: []interface{}{10, 5, 15, 3, 7, nil, 18},
				low:  7, high: 15},
			Output: 32,
		},
		{
			Input: Input{
				root: []interface{}{10, 5, 15, 3, 7, 13, 18, 1, nil, 6},
				low:  6, high: 10},
			Output: 23,
		},
	}
}

func TestRangeSumBST(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input.root)
		output := rangeSumBST(tree.Root, testCase.Input.low, testCase.Input.high)
		assert.Equal(t, testCase.Output, output)
	}
}
