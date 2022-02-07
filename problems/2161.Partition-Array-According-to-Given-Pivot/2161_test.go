package p2161

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums  []int
	pivot int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{9, 12, 5, 10, 14, 3, 10}, 10}, []int{9, 5, 3, 10, 10, 12, 14}},
		{Input{[]int{-3, 4, 3, 2}, 2}, []int{-3, 2, 4, 3}},
	}
}

func TestPivotArray(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := pivotArray(testCase.Input.nums, testCase.Input.pivot)
		assert.Equal(t, testCase.Output, output)
	}
}
