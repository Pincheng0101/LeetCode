package p1365

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
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}
}

func TestSmallerNumbersThanCurrent_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := smallerNumbersThanCurrent_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestSmallerNumbersThanCurrent_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := smallerNumbersThanCurrent_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
