package p2154

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums     []int
	original int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{5, 3, 6, 1, 12}, 3}, 24},
		{Input{[]int{2, 7, 9}, 4}, 4},
	}
}

func TestFindFinalValue_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findFinalValue_1(testCase.Input.nums, testCase.Input.original)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestFindFinalValue_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findFinalValue_2(testCase.Input.nums, testCase.Input.original)
		assert.Equal(t, testCase.Output, output)
	}
}
