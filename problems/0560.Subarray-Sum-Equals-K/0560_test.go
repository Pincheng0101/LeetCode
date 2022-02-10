package p0560

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
	k    int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 1, 1}, 2}, 2},
		{Input{[]int{1, 2, 3}, 3}, 2},
	}
}

func TestSubarraySum_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := subarraySum_1(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestSubarraySum_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := subarraySum_2(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
