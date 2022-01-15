package p1342

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
		{14, 6},
		{8, 4},
		{123, 12},
	}
}

func TestNumberOfSteps(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numberOfSteps(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
