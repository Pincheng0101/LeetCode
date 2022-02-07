package p2165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  int64
	Output int64
}

func testCases() []testCase {
	return []testCase{
		{310, 103},
		{-7605, -7650},
	}
}

func TestSmallestNumber(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := smallestNumber(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
