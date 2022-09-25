package p2419

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
		{[]int{1, 2, 3, 3, 2, 2}, 2},
		{[]int{1, 2, 3, 4}, 1},
	}
}

func TestLongestSubarray(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := longestSubarray(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
