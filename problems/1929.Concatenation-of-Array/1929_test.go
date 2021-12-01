package p1929

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
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}
}

func TestGetConcatenation(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := getConcatenation(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
