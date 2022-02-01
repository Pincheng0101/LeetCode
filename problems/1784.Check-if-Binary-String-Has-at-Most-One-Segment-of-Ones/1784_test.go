package p1784

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{"1001", false},
		{"110", true},
	}
}

func TestCheckOnesSegment(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkOnesSegment(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
