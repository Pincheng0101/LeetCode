package p0525

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
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1}, 14},
	}
}

func TestFindMaxLength(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findMaxLength(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
