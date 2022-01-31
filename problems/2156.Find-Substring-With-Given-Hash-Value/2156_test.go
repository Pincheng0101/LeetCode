package p2156

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s         string
	power     int
	modulo    int
	k         int
	hashValue int
}

type testCase struct {
	Input  Input
	Output string
}

func testCases() []testCase {
	return []testCase{
		{Input{"leetcode", 7, 20, 2, 0}, "ee"},
		{Input{"fbxzaad", 31, 100, 3, 32}, "fbx"},
		{Input{"xmmhdakfursinye", 96, 45, 15, 21}, "xmmhdakfursinye"},
		{Input{"xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk", 22, 51, 41, 9}, "xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxti"},
	}
}

func TestSubStrHash(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := subStrHash(testCase.Input.s, testCase.Input.power, testCase.Input.modulo, testCase.Input.k, testCase.Input.hashValue)
		assert.Equal(t, testCase.Output, output)
	}
}
