package p0680

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{"aba", true},
		{"abca", true},
		{"abc", false},
	}
}

func TestValidPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := validPalindrome(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
