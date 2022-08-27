package p1162

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
		{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, 2},
		{[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 4},
	}
}

func TestMaxDistance(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxDistance(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
