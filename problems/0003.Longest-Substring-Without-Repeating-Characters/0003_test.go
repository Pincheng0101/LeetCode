package p0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"tmmzuxt", 5},
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := lengthOfLongestSubstring(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
