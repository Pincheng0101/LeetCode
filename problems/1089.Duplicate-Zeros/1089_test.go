package p1089

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
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}
}

func TestDuplicateZeros_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		duplicateZeros_1(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}
