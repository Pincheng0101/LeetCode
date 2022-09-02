package p0542

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
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}
}

func TestUpdateMatrix(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := updateMatrix(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
