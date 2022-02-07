package p2163

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
		{[]int{3, 1, 2}, -1},
		{[]int{7, 9, 5, 8, 1, 3}, 1},
	}
}

func TestMinimumDifference(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minimumDifference(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
