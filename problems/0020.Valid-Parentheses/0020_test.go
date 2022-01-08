package p0020

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
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"){}", false},
	}
}

func TestIsValid(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isValid(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
