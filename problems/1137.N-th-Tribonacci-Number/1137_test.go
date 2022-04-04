package p1137

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
		{4, 4},
		{25, 1389537},
	}
}

func TestTribonacci_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := tribonacci_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestTribonacci_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := tribonacci_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
