package p0127

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	beginWord string
	endWord   string
	wordList  []string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}}, 5},
		{Input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}}, 0},
	}
}

func TestLadderLength(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := ladderLength(testCase.Input.beginWord, testCase.Input.endWord, testCase.Input.wordList)
		assert.Equal(t, testCase.Output, output)
	}
}
