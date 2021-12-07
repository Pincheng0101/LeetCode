package p1108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output string
}

func testCases() []testCase {
	return []testCase{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}
}

func TestDefangIPaddr(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := defangIPaddr(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
