package p2420

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
		{Input{[]int{2, 1, 1, 1, 3, 4, 1}, 2}, []int{2, 3}},
		{Input{[]int{2, 1, 1, 2}, 2}, []int{}},
	}
}

func TestGoodIndices(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := goodIndices(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
