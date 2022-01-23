package p2151

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
		{[][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}, 2},
		{[][]int{{2, 0}, {0, 2}}, 1},
	}
}

func TestMaximumGood(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maximumGood(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
