package p0066

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
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}
}

func TestPlusOne(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := plusOne(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
