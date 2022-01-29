package p0045

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1}, 0},
		{[]int{1, 1, 1, 1}, 3},
	}
}

func TestJump_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := jump_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestJump_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := jump_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
