package p0152

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
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}
}

func TestMaxProduct(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxProduct(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
