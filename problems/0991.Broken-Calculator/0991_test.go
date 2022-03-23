package p0991

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	startValue int
	target     int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{2, 3}, 2},
		{Input{5, 8}, 2},
		{Input{3, 10}, 3},
	}
}

func TestBrokenCalc(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := brokenCalc(testCase.Input.startValue, testCase.Input.target)
		assert.Equal(t, testCase.Output, output)
	}
}
