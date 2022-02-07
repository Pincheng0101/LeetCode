package p2164

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
		{[]int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{[]int{2, 1}, []int{2, 1}},
	}
}

func TestSortEvenOdd(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sortEvenOdd(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
