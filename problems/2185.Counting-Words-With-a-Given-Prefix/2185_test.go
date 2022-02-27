package p2185

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	words []string
	pref  string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]string{"pay", "attention", "practice", "attend"}, "at"}, 2},
		{Input{[]string{"leetcode", "win", "loops", "success"}, "code"}, 0},
	}
}

func TestPrefixCount(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := prefixCount(testCase.Input.words, testCase.Input.pref)
		assert.Equal(t, testCase.Output, output)
	}
}
