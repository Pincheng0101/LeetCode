package p2180

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
		{4, 2},
		{30, 14},
	}
}

func TestCountEven(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := countEven(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
