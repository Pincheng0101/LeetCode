package p0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Nums   []int
	Target int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{
			Input:  Input{Nums: []int{2, 7, 11, 15}, Target: 9},
			Output: []int{0, 1},
		},
		{
			Input:  Input{Nums: []int{3, 2, 4}, Target: 6},
			Output: []int{1, 2},
		},
		{
			Input:  Input{Nums: []int{3, 3}, Target: 6},
			Output: []int{0, 1},
		},
	}
}

func TestTwoSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := twoSum(testCase.Input.Nums, testCase.Input.Target)
		assert.Equal(t, testCase.Output, output)
	}
}
