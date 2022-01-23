package p0520

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
		{"USA", true},
		{"FlaG", false},
		{"FFaa", false},
	}
}

func TestDetectCapitalUse(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := detectCapitalUse(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
