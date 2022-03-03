package p0413

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
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1}, 0},
		{[]int{1, 2, 3, 4, 6, 8, 10, 12}, 9},
	}
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numberOfArithmeticSlices(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
