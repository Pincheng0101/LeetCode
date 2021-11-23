package p0941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
	}
}

func TestValidMountainArray(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := validMountainArray(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
