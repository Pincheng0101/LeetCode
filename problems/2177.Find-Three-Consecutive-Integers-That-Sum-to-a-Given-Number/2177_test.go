package p2177

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
		{33, []int64{10, 11, 12}},
		{4, []int64{}},
	}
}

func TestSumOfThree(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sumOfThree(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
