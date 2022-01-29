package p0561

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
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}
}

func TestArrayPairSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := arrayPairSum(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
