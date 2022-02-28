package p0228

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output []string
}

func testCases() []testCase {
	return []testCase{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{-1}, []string{"-1"}},
	}
}

func TestSummaryRanges(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := summaryRanges(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
