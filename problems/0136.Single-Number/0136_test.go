package p0136

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
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
}

func TestSingleNumber(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := singleNumber(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
