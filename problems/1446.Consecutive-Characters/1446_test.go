package p1446

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{"", 0},
		{"leetcode", 2},
		{"abbcccddddeeeeedcba", 5},
		{"triplepillooooow", 5},
		{"hooraaaaaaaaaaay", 11},
		{"tourist", 1},
	}
}

func TestMaxPower(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := maxPower(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
