package p1091

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  [][]int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
	}
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := shortestPathBinaryMatrix(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
