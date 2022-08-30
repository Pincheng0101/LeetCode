package p0417

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
		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{
				{0, 4}, {1, 3},
				{1, 4}, {2, 2},
				{3, 0}, {3, 1},
				{4, 0}},
		},
		{
			[][]int{
				{1},
			},
			[][]int{
				{0, 0},
			},
		},
	}
}

func TestPacificAtlantic(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := pacificAtlantic(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
