package p2233

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
		{Input{[]int{0, 4}, 5}, 20},
		{Input{[]int{6, 3, 3, 2}, 2}, 216},
	}
}

func TestMaximumProduct(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maximumProduct(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
