package p0740

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
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
	}
}

func TestDeleteAndEarn(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := deleteAndEarn(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
