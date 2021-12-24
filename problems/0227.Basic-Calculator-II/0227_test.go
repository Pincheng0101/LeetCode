package p0227

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{"3+5 / 2", 5},
		{"1 + 1", 2},
	}
}

func TestCalculate(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := calculate(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
