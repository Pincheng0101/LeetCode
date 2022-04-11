package p0682

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
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{[]string{"1"}, 1},
	}
}

func TestCalPoints(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := calPoints(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
