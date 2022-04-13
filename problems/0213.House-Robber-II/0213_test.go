package p0213

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
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
		{[]int{1}, 1},
	}
}

func TestRob(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := rob(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
