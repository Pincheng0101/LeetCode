package p0488

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
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
	}
}

func TestFindDisappearedNumbers_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findDisappearedNumbers_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestFindDisappearedNumbers_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findDisappearedNumbers_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
