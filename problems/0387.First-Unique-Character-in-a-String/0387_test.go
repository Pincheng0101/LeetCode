package p0387

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
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
}

func TestFirstUniqChar(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := firstUniqChar(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
