package p2400

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	startPos int
	endPos   int
	k        int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{1, 2, 3}, 3},
		{Input{2, 5, 10}, 0},
		{Input{1000, 989, 999}, 903199077},
	}
}

func TestNumberOfWays(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numberOfWays(testCase.Input.startPos, testCase.Input.endPos, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
