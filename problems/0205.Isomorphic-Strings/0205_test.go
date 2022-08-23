package p0205

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
		{Input{"egg", "add"}, true},
		{Input{"foo", "bar"}, false},
		{Input{"paper", "title"}, true},
		{Input{"badc", "baba"}, false},
	}
}

func TestIsIsomorphic(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isIsomorphic(testCase.Input.s, testCase.Input.t)
		assert.Equal(t, testCase.Output, output)
	}
}
