package p0287

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
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}
}

func TestFindDuplicate(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findDuplicate(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
