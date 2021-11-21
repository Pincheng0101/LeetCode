package p0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Nums []int
	Val  int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{
			Input: Input{
				Nums: []int{3, 2, 2, 3}, Val: 3},
			Output: 2,
		},
		{
			Input: Input{
				Nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, Val: 2},
			Output: 5,
		},
	}
}

func TestRemoveElement(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := removeElement(testCase.Input.Nums, testCase.Input.Val)
		assert.Equal(t, testCase.Output, output)
	}
}
