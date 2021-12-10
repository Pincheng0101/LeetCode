package p1672

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
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}
}

func TestMaximumWealth(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maximumWealth(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
