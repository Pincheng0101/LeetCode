package p0338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
}

func TestCountBits(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := countBits(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
