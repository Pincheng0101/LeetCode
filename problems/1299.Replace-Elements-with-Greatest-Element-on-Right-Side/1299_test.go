package p1299

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
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{[]int{400}, []int{-1}},
	}
}

func TestReplaceElements(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		replaceElements(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}
