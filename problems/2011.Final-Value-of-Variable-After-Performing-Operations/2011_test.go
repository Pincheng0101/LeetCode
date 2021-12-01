package p2011

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
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}
}

func TestFinalValueAfterOperations(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := finalValueAfterOperations(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
