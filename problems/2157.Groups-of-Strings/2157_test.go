package p2157

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []string
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]string{"a", "b", "ab", "cde"}, []int{2, 3}},
		{[]string{"a", "ab", "abc"}, []int{1, 3}},
	}
}

func TestGroupStrings(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := groupStrings(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
