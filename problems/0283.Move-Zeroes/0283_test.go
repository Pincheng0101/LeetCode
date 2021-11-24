package p0283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{
			[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0},
		},
		{
			[]int{0}, []int{0},
		},
	}
}

func TestMoveZeroes(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		moveZeroes(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}
