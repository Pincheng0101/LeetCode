package p1663

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	n int
	k int
}

type testCase struct {
	Input  Input
	Output string
}

func testCases() []testCase {
	return []testCase{
		{Input{3, 27}, "aay"},
		{Input{5, 73}, "aaszz"},
	}
}

func TestGetSmallestString(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := getSmallestString(testCase.Input.n, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
