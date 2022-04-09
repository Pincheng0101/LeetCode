package p0347

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
		{Input{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{Input{[]int{1}, 1}, []int{1}},
	}
}

func TestTopKFrequent(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := topKFrequent(testCase.Input.nums, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
