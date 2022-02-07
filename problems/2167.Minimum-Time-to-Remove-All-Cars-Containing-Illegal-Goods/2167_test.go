package p2167

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
		{"1100101", 5},
		{"0010", 2},
	}
}

func TestMinimumTime(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minimumTime(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
