package p0198

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
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
}

func TestRob(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := rob(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
