package p2399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s        string
	distance []int
}

type testCase struct {
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, true},
		{Input{"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, false},
	}
}

func TestCheckDistances(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkDistances(testCase.Input.s, testCase.Input.distance)
		assert.Equal(t, testCase.Output, output)
	}
}
