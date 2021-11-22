package p0026

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
		{
			Input:  []int{1, 1, 2},
			Output: 2,
		},
		{
			Input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Output: 5,
		},
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := removeDuplicates(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
