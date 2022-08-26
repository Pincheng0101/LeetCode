package p1020

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
		{[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, 3},
		{[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, 0},
	}
}

func TestNumEnclaves(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numEnclaves(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
