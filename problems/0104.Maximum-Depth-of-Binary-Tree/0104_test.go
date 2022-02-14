package p0104

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/binarytree"
)

type testCase struct {
	Input  []interface{}
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, 3},
		{[]interface{}{1, nil, 2}, 2},
	}
}

func TestMaxDepth(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := maxDepth(tree.Root)
		assert.Equal(t, testCase.Output, output)
	}
}
