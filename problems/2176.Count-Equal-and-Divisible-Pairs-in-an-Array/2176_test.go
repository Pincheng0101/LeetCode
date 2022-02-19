package p2176

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
	k    int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{3, 1, 2, 2, 2, 1, 3}, 2}, 4},
		{Input{[]int{1, 2, 3, 4}, 1}, 0},
	}
}

func TestCountPairs(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := countPairs(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
