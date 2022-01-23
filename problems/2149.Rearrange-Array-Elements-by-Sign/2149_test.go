package p2149

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]int{3, 1, -2, -5, 2, -4}, []int{3, -2, 1, -5, 2, -4}},
		{[]int{-1, 1}, []int{1, -1}},
		{[]int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}, []int{28, -41, 22, -8, 46, -37, 35, -9, 18, -6, 19, -26, 15, -37, 14, -10, 31, -9}},
	}
}

func TestRearrangeArray_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := rearrangeArray_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestRearrangeArray_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := rearrangeArray_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
