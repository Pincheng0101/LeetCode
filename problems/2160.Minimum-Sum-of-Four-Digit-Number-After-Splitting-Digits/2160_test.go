package p2160

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
		{2932, 52},
		{4009, 13},
	}
}

func TestMinimumSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minimumSum(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
