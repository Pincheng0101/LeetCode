package p0189

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
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{Input{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}
}

func TestRotate_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		rotate_1(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, testCase.Input.nums)
	}
}

func TestRotate_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		rotate_2(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, testCase.Input.nums)
	}
}
