package p0454

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums1 []int
	nums2 []int
	nums3 []int
	nums4 []int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}}, 2},
		{Input{[]int{0}, []int{0}, []int{0}, []int{0}}, 1},
	}
}

func TestFourSumCount(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := fourSumCount(testCase.Input.nums1, testCase.Input.nums2, testCase.Input.nums3, testCase.Input.nums4)
		assert.Equal(t, testCase.Output, output)
	}
}
