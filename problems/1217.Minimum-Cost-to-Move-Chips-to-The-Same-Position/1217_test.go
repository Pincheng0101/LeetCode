package p1217

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
		{[]int{1, 2, 3}, 1},
		{[]int{2, 2, 2, 3, 3}, 2},
		{[]int{1, 1000000000}, 1},
	}
}

func TestMinCostToMoveChips(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minCostToMoveChips(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
