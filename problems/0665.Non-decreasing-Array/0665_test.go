package p0665

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
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{-100000, -100000, -100000, 100000}, true},
		{[]int{1, 1, 1}, true},
	}
}

func TestCheckPossibility(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkPossibility(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
