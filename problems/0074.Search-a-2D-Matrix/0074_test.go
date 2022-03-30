package p0074

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	matrix [][]int
	target int
}

type testCase struct {
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3}, true},
		{Input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13}, false},
	}
}

func TestSearchMatrix(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := searchMatrix(testCase.Input.matrix, testCase.Input.target)
		assert.Equal(t, testCase.Output, output)
	}
}
