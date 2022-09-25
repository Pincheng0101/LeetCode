package p0700

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/binarytree"
)

type Input struct {
	root []interface{}
	val  int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]interface{}{4, 2, 7, 1, 3}, 2}, []int{2, 1, 3}},
		{Input{[]interface{}{4, 2, 7, 1, 3}, 5}, []int{}},
	}
}

func TestSearchBST(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		tree := NewBinaryTree(testCase.Input.root)
		output := searchBST(tree.Root, testCase.Input.val)
		assert.Equal(t, testCase.Output, output.GetNodeValuesWithDFS())
	}
}
