package p1010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{30, 20, 150, 100, 40}, 3},
		{[]int{60, 60, 60}, 3},
		{[]int{15, 63, 451, 213, 37, 209, 343, 319}, 1},
	}
}

func TestNumPairsDivisibleBy60(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numPairsDivisibleBy60(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
