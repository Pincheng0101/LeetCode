package p1512

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
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}
}

func TestNumIdenticalPairs(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numIdenticalPairs(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
