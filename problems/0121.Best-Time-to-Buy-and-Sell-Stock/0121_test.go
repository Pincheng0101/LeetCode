package p0121

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
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
}

func TestMaxProfit(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxProfit(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
