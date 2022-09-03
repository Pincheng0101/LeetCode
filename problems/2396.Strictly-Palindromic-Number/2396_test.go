package p2396

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
		{9, false},
		{4, false},
	}
}

func TestIsStrictlyPalindromic(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isStrictlyPalindromic(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
