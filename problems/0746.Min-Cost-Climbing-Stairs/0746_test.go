package p0746

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
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minCostClimbingStairs(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
