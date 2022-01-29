package p0409

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
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
		{"ccc", 3},
	}
}

func TestLongestPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := longestPalindrome(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
