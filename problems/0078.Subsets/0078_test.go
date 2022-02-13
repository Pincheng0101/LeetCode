package p0078

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
}

func TestSubsets(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := subsets(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
