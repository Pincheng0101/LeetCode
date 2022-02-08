package p0258

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
		{38, 2},
		{0, 0},
	}
}

func TestAddDigits(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := addDigits(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
