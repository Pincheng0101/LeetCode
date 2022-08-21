package p2381

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s      string
	shifts [][]int
}

type testCase struct {
	Input  Input
	Output string
}

func testCases() []testCase {
	return []testCase{
		{Input{"abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}}, "ace"},
		{Input{"dztz", [][]int{{0, 0, 0}, {1, 1, 1}}}, "catz"},
	}
}

func TestShiftingLetters(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := shiftingLetters(testCase.Input.s, testCase.Input.shifts)
		assert.Equal(t, testCase.Output, output)
	}
}
