package p0438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s string
	p string
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{"cbaebabacd", "abc"}, []int{0, 6}},
		{Input{"abab", "ab"}, []int{0, 1, 2}},
	}
}

func TestFindAnagrams(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findAnagrams(testCase.Input.s, testCase.Input.p)
		assert.Equal(t, testCase.Output, output)
	}
}
