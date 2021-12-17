package p1476

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubrectangleQueries_1(t *testing.T) {
	subrectangleQueries := SubrectangleQueries{[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}}
	assert.Equal(t, 1, subrectangleQueries.GetValue(0, 2))
	subrectangleQueries.UpdateSubrectangle(0, 0, 3, 2, 5)
	assert.Equal(t, 5, subrectangleQueries.GetValue(0, 2))
	assert.Equal(t, 5, subrectangleQueries.GetValue(3, 1))
	subrectangleQueries.UpdateSubrectangle(3, 0, 3, 2, 10)
	assert.Equal(t, 10, subrectangleQueries.GetValue(3, 1))
	assert.Equal(t, 5, subrectangleQueries.GetValue(0, 2))
}

func TestSubrectangleQueries_2(t *testing.T) {
	subrectangleQueries := SubrectangleQueries{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}}
	assert.Equal(t, 1, subrectangleQueries.GetValue(0, 0))
	subrectangleQueries.UpdateSubrectangle(0, 0, 2, 2, 100)
	assert.Equal(t, 100, subrectangleQueries.GetValue(0, 0))
	assert.Equal(t, 100, subrectangleQueries.GetValue(2, 2))
	subrectangleQueries.UpdateSubrectangle(1, 1, 2, 2, 20)
	assert.Equal(t, 20, subrectangleQueries.GetValue(2, 2))
}
