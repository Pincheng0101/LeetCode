package p0977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
}

func TestSortedSquares_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sortedSquares_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestSortedSquares_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sortedSquares_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
