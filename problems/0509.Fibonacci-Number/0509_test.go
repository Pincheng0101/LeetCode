package p0509

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
		{2, 1},
		{1000, 817770325994397771},
	}
}

func TestFib_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := fib_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestFib_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := fib_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
