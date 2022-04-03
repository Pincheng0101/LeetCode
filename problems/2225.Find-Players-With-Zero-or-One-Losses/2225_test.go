package p2225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  [][]int
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}, [][]int{{1, 2, 10}, {4, 5, 7, 8}}},
		{[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}, [][]int{{1, 2, 5, 6}, {}}},
	}
}

func TestFindWinners(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findWinners(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
