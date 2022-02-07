package p0389

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
	Output byte
}

func testCases() []testCase {
	return []testCase{
		{Input{"abcd", "abcde"}, byte('e')},
		{Input{"", "y"}, byte('y')},
	}
}

func TestFindTheDifference(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findTheDifference(testCase.Input.s, testCase.Input.t)
		assert.Equal(t, testCase.Output, output)
	}
}
