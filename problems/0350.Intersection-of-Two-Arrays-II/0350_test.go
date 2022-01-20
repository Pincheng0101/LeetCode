package p0350

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
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{Input{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{9, 4}},
	}
}

func TestIntersect(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := intersect(testCase.Input.nums1, testCase.Input.nums2)
		assert.Equal(t, testCase.Output, output)
	}
}
