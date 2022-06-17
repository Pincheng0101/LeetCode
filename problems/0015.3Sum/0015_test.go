package p0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
	}
}

func TestThreeSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := threeSum(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
