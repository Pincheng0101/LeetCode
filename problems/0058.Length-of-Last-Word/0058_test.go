package p0058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
}

func TestLengthOfLastWord(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := lengthOfLastWord(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
