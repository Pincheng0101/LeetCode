package p2162

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	startAt       int
	moveCost      int
	pushCost      int
	targetSeconds int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{1, 2, 1, 600}, 6},
		{Input{0, 1, 4, 9}, 5},
		{Input{0, 1, 1, 1}, 2},
		{Input{0, 1, 2, 76}, 6},
		{Input{0, 6578, 6577, 17}, 26310},
		{Input{0, 1, 1, 6039}, 5},
	}
}

func TestMinCostSetTime(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minCostSetTime(testCase.Input.startAt, testCase.Input.moveCost, testCase.Input.pushCost, testCase.Input.targetSeconds)
		assert.Equal(t, testCase.Output, output)
	}
}
