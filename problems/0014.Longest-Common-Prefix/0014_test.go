package p0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []string
	Output string
}

func testCases() []testCase {
	return []testCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := longestCommonPrefix(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
