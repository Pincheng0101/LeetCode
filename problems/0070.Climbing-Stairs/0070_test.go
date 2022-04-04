package p0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{2, 2},
		{45, 1836311903},
	}
}

func TestClimbStairs(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := climbStairs(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
