package p0011

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
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
}

func TestMaxArea(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxArea(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
