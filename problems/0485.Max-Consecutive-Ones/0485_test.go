package p0485

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findMaxConsecutiveOnes(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
