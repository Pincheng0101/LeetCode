package p2166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunction(t *testing.T) {
	bitset := Constructor(2)
	bitset.Flip()
	bitset.Unfix(1)
	assert.Equal(t, bitset.All(), false)
	bitset.Fix(1)
	bitset.Fix(1)
	bitset.Unfix(1)
	assert.Equal(t, bitset.All(), false)
	assert.Equal(t, bitset.Count(), 1)
	assert.Equal(t, bitset.ToString(), "10")
	bitset.Unfix(0)
	bitset.Flip()
	bitset.Unfix(0)
	assert.Equal(t, bitset.All(), false)
	assert.Equal(t, bitset.One(), true)
}
