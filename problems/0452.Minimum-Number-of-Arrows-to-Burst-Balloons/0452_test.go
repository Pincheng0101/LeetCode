package p0452

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
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}
}

func TestFindMinArrowShots_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findMinArrowShots_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestFindMinArrowShots_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findMinArrowShots_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
