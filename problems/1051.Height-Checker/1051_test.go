package p1051

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
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
}

func TestHeightChecker(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := heightChecker(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
