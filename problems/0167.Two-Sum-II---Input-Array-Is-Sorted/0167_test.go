package p0167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	numbers []int
	target  int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{Input{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{Input{[]int{-1, 0}, -1}, []int{1, 2}},
	}
}

func TestTwoSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := twoSum(testCase.Input.numbers, testCase.Input.target)
		assert.Equal(t, testCase.Output, output)
	}
}
