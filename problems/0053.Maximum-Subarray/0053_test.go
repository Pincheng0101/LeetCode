package p0053

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
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
}

func TestMaxSubArray(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxSubArray(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
