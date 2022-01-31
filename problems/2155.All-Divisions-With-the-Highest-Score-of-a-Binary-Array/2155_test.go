package p1001

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
		{[]int{0, 0, 1, 0}, []int{2, 4}},
		{[]int{0, 0, 0}, []int{3}},
		{[]int{1, 1}, []int{0}},
	}
}

func TestMaxScoreIndices(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxScoreIndices(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
