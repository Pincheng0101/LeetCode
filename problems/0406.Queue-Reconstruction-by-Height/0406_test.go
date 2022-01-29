package p0406

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  [][]int
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}, [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
		{[][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}}, [][]int{{3, 0}, {6, 0}, {7, 0}, {5, 2}, {3, 4}, {5, 3}, {6, 2}, {2, 7}, {9, 0}, {1, 9}}},
	}
}

func TestReconstructQueue_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := reconstructQueue_1(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestReconstructQueue_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := reconstructQueue_2(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestReconstructQueue_3(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := reconstructQueue_3(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}

func BenchmarkReconstructQueue_LinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCases := testCases()
		for _, testCase := range testCases {
			reconstructQueue_1(testCase.Input)
		}
	}
}

func BenchmarkReconstructQueue_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCases := testCases()
		for _, testCase := range testCases {
			reconstructQueue_2(testCase.Input)
		}
	}
}

func BenchmarkReconstructQueue_Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCases := testCases()
		for _, testCase := range testCases {
			reconstructQueue_3(testCase.Input)
		}
	}
}
