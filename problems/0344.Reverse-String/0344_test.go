package p0344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []byte
	Output []byte
}

func testCases() []testCase {
	return []testCase{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
}

func TestReverseString(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		reverseString(testCase.Input)
		assert.Equal(t, testCase.Output, testCase.Input)
	}
}
