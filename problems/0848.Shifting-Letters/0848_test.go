package p0848

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s      string
	shifts []int
}

type testCase struct {
	Input  Input
	Output string
}

func testCases() []testCase {
	return []testCase{
		{Input{"abc", []int{3, 5, 9}}, "rpl"},
		{Input{"aaa", []int{1, 2, 3}}, "gfd"},
	}
}

func TestShiftingLetters(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := shiftingLetters(testCase.Input.s, testCase.Input.shifts)
		assert.Equal(t, testCase.Output, output)
	}
}
