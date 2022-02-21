package p0169

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
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
}

func TestMajorityElement(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := majorityElement(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
