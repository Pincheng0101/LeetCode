package p0414

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
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
	}
}

func TestThirdMax_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := thirdMax_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestThirdMax_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := thirdMax_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
