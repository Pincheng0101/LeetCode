package p2413

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
		{5, 10},
		{6, 6},
	}
}

func TestFunction(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := smallestEvenMultiple(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
