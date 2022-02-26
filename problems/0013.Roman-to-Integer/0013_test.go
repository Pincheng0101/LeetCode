package p0013

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
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
}

func TestRomanToInt(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := romanToInt(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
