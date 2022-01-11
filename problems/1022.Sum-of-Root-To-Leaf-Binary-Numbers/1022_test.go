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

func TestSumRootToLeaf_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := sumRootToLeaf_1(tree.Root)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestSumRootToLeaf_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input)
		output := sumRootToLeaf_2(tree.Root)
		assert.Equal(t, testCase.Output, output)
	}
}
