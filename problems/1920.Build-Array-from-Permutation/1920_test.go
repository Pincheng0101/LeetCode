package p1920

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
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}
}

func TestGetConcatenation(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := buildArray(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
