package p2221

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
		{[]int{1, 2, 3, 4, 5}, 8},
	}
}

func TestTriangularSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := triangularSum(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
