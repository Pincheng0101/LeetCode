package p0605

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	flowerbed []int
	n         int
}

type testCase struct {
	Input  Input
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 0, 0, 0, 1}, 1}, true},
		{Input{[]int{1, 0, 0, 0, 1}, 2}, false},
	}
}

func TestCanPlaceFlowers(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := canPlaceFlowers(testCase.Input.flowerbed, testCase.Input.n)
		assert.Equal(t, testCase.Output, output)
	}
}
