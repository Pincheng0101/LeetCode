package p2416

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
		{[]string{"abc", "ab", "bc", "b"}, []int{5, 4, 3, 2}},
		{[]string{"abcd"}, []int{4}},
	}
}

func TestSumPrefixScores(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sumPrefixScores(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
