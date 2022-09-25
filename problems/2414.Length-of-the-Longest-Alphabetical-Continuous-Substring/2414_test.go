package p2414

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
		{"abacaba", 2}, {"abcde", 5},
	}
}

func TestLongestContinuousSubstring(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := longestContinuousSubstring(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
