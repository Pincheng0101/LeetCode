package p1288

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
		{[][]int{{1, 4}, {3, 6}, {2, 8}}, 2},
		{[][]int{{1, 4}, {2, 3}}, 1},
	}
}

func TestRemoveCoveredIntervals(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := removeCoveredIntervals(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
