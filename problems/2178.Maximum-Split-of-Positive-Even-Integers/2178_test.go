package p2178

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  int64
	Output []int64
}

func testCases() []testCase {
	return []testCase{
		{12, []int64{2, 4, 6}},
		{7, []int64{}},
		{28, []int64{2, 4, 6, 16}},
	}
}

func TestMaximumEvenSplit(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maximumEvenSplit(testCase.Input)
		assert.ElementsMatch(t, testCase.Output, output)
	}
}
