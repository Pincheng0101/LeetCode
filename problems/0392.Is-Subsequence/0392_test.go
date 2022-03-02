package p0392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s string
	t string
}

type testCase struct {
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{"abc", "ahbgdc"}, true},
		{Input{"axc", "ahbgdc"}, false},
	}
}

func TestIsSubsequence(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isSubsequence(testCase.Input.s, testCase.Input.t)
		assert.Equal(t, testCase.Output, output)
	}
}
