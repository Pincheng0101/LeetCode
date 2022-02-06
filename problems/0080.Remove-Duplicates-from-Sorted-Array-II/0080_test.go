package p0080

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Output struct {
	expectedNums []int
	k            int
}

type testCase struct {
	Input  []int
	Output Output
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 1, 1, 2, 2, 3}, Output{[]int{1, 1, 2, 2, 3}, 5}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, Output{[]int{0, 0, 1, 1, 2, 3, 3}, 7}},
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := removeDuplicates(testCase.Input)
		assert.Equal(t, testCase.Output.k, output)
		assert.Equal(t, testCase.Output.expectedNums, testCase.Input[:output])
	}
}
