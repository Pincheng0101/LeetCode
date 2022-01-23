package p2150

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
		{[]int{10, 6, 5, 8}, []int{10, 8}},
		{[]int{1, 3, 5, 3}, []int{1, 5}},
	}
}

func TestFindLonely_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findLonely_1(testCase.Input)
		assert.ElementsMatch(t, testCase.Output, output)
	}
}

func TestFindLonely_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findLonely_2(testCase.Input)
		assert.ElementsMatch(t, testCase.Output, output)
	}
}
