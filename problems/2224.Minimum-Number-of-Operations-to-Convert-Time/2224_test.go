package p2224

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	current string
	correct string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{"02:30", "04:35"}, 3},
		{Input{"11:00", "11:01"}, 1},
	}
}

func TestConvertTime(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := convertTime(testCase.Input.current, testCase.Input.correct)
		assert.Equal(t, testCase.Output, output)
	}
}
