package p2222

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int64
}

func testCases() []testCase {
	return []testCase{
		{"001101", 6},
		{"11100", 0},
	}
}

func TestNumberOfWays(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numberOfWays(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
