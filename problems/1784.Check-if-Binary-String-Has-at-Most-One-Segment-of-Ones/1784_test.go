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

func TestCheckOnesSegment_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkOnesSegment_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestCheckOnesSegment_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkOnesSegment_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
