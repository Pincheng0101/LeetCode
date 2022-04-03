package p2220

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	start int
	goal  int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{10, 7}, 3},
		{Input{3, 4}, 3},
	}
}

func TestMinBitFlips(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minBitFlips(testCase.Input.start, testCase.Input.goal)
		assert.Equal(t, testCase.Output, output)
	}
}
