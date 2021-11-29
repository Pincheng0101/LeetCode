package p0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Haystack string
	Needle   string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{
			Input: Input{
				Haystack: "hello", Needle: "ll"},
			Output: 2,
		},
		{
			Input: Input{
				Haystack: "aaaaa", Needle: "bba"},
			Output: -1,
		},
		{
			Input: Input{
				Haystack: "", Needle: ""},
			Output: 0,
		},
		{
			Input: Input{
				Haystack: "abc", Needle: "c"},
			Output: 2,
		},
	}
}

func TestStrStr_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := strStr_1(testCase.Input.Haystack, testCase.Input.Needle)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestStrStr_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := strStr_2(testCase.Input.Haystack, testCase.Input.Needle)
		assert.Equal(t, testCase.Output, output)
	}
}
