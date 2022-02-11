package p0567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s1 string
	s2 string
}

type testCase struct {
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{"ab", "eidbaooo"}, true},
		{Input{"ab", "eidboaoo"}, false},
	}
}

func TestCheckInclusion(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkInclusion(testCase.Input.s1, testCase.Input.s2)
		assert.Equal(t, testCase.Output, output)
	}
}
