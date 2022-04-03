package p2226

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	candies []int
	k       int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{5, 8, 6}, 3}, 5},
		{Input{[]int{2, 5}, 11}, 0},
	}
}

func TestMaximumCandies(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maximumCandies(testCase.Input.candies, int64(testCase.Input.k))
		assert.Equal(t, testCase.Output, output)
	}
}
