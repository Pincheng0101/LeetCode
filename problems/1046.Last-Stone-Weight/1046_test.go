package p1046

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
	}
}

func TestLastStoneWeight(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := lastStoneWeight(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
