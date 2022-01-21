package p0125

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
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isPalindrome(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
