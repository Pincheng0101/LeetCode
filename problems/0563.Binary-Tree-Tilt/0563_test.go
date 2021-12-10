package p0563

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/binarytree"
)

type root []interface{}

type testCase struct {
	Input  root
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]interface{}{1, 2, 3}, 1},
		{[]interface{}{4, 2, 9, 3, 5, nil, 7}, 15},
	}
}

func TestFindTilt(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := findTilt(tree.Root)
		assert.Equal(t, testCase.Output, output)
	}
}
