package p0007

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
		{123, 321},
		{-123, -321},
		{120, 21},
	}
}

func TestReverse(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := reverse(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
