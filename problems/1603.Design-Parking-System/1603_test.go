package p1603

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParkingSystem(t *testing.T) {
	parkingSystem := Constructor(1, 1, 0)
	assert.Equal(t, parkingSystem.AddCar(1), true)
	assert.Equal(t, parkingSystem.AddCar(2), true)
	assert.Equal(t, parkingSystem.AddCar(3), false)
}
