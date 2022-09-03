package p2395

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
		{[]int{4, 2, 4}, true},
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{0, 0, 0}, true},
	}
}

func TestFindSubarrays(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findSubarrays(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
