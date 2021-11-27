package p0217

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
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
}

func TestContainsDuplicate(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := containsDuplicate(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
