package p0135

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
}

func TestCandy(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := candy(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
