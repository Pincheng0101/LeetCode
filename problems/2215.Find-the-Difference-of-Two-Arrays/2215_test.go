package p2215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums1 []int
	nums2 []int
}

type testCase struct {
	Input  Input
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3}, []int{2, 4, 6}}, [][]int{{1, 3}, {4, 6}}},
		{Input{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}}, [][]int{{3}, {}}},
	}
}

func TestFindDifference(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findDifference(testCase.Input.nums1, testCase.Input.nums2)
		assert.Equal(t, testCase.Output, output)
	}
}
