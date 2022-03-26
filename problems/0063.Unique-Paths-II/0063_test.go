package p0063

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
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := uniquePathsWithObstacles(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
