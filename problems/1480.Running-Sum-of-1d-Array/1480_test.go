package p1480

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
}

func TestRunningSum(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := runningSum(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
