package p0724

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
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}
}

func TestPivotIndex(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := pivotIndex(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
