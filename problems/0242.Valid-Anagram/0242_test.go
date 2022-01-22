package p0242

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
		{Input{"anagram", "nagaram"}, true},
		{Input{"rat", "car"}, false},
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := isAnagram(testCase.Input.s, testCase.Input.t)
		assert.Equal(t, testCase.Output, output)
	}
}
