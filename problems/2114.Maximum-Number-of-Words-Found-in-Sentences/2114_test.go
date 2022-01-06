package p2114

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	}
}

func TestMostWordsFound(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := mostWordsFound(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
