package p1295

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
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
}

func TestFindNumbers_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findNumbers_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestFindNumbers_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findNumbers_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
