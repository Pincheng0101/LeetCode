package p0532

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
		{Input{[]int{3, 1, 4, 1, 5}, 2}, 2},
		{Input{[]int{1, 2, 3, 4, 5}, 1}, 4},
		{Input{[]int{1, 3, 1, 5, 4}, 0}, 1},
	}
}

func TestFindPairs(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findPairs(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
