package p1089

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
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 1, 7, 6, 0, 2, 0, 7}, []int{0, 0, 1, 7, 6, 0, 0, 2}},
		{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
		{[]int{1, 5, 2, 0, 6, 8, 0, 6, 0}, []int{1, 5, 2, 0, 0, 6, 8, 0, 0}},
	}
}

func TestDuplicateZeros_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		duplicateZeros_1(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}

func TestDuplicateZeros_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		duplicateZeros_2(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}

func TestDuplicateZeros_3(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		duplicateZeros_3(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}
