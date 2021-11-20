package p0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Nums1 []int
	M     int
	Nums2 []int
	N     int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{
			Input: Input{
				Nums1: []int{1, 2, 3, 0, 0, 0}, M: 3,
				Nums2: []int{2, 5, 6}, N: 3},
			Output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Input: Input{
				Nums1: []int{1}, M: 1,
				Nums2: []int{}, N: 0},
			Output: []int{1},
		},
		{
			Input: Input{
				Nums1: []int{0}, M: 0,
				Nums2: []int{1}, N: 1},
			Output: []int{1},
		},
		{
			Input: Input{
				Nums1: []int{4, 5, 6, 0, 0, 0}, M: 3,
				Nums2: []int{1, 2, 3}, N: 3},
			Output: []int{1, 2, 3, 4, 5, 6},
		},
	}
}

func TestMerge1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		merge1(testCase.Input.Nums1, testCase.Input.M, testCase.Input.Nums2, testCase.Input.N)
		assert.Equal(t, testCase.Output, testCase.Input.Nums1)
	}
}

func TestMerge2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		merge2(testCase.Input.Nums1, testCase.Input.M, testCase.Input.Nums2, testCase.Input.N)
		assert.Equal(t, testCase.Output, testCase.Input.Nums1)
	}
}
