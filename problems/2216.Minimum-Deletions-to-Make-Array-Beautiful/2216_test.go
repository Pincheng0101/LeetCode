package p2216

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
		{[]int{1, 1, 2, 3, 5}, 1},
		{[]int{1, 1, 2, 2, 3, 3}, 2},
	}
}

func TestMinDeletion(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minDeletion(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
