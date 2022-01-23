package p2148

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
		{[]int{11, 7, 2, 15}, 2},
		{[]int{-3, 3, 3, 90}, 2},
	}
}

func TestCountElements(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := countElements(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
