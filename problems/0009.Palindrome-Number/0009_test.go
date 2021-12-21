package p0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  int
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{121, true},
		{-121, false},
		{10, false},
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isPalindrome(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
