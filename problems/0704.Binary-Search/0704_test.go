package p0704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums   []int
	target int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{Input{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}
}

func TestSearch(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := search(testCase.Input.nums, testCase.Input.target)
		assert.Equal(t, testCase.Output, output)
	}
}
