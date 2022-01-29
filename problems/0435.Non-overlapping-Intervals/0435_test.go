package p0435

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
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	}
}

func TestEraseOverlapIntervals(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := eraseOverlapIntervals(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
