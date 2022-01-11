package p1022

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
		{[]interface{}{1, 0, 1, 0, 1, 0, 1}, 22},
		{[]interface{}{0}, 0},
	}
}

func TestSumRootToLeaf(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := sumRootToLeaf(tree.Root)
		assert.Equal(t, testCase.Output, output)
	}
}
