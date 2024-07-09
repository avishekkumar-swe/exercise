package parkinglot_test

import (
	parkinglot "github.com/avishekkumar-swe/exercise/tree/master"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingLot(t *testing.T) {
	t.Run("should be able to park a car", func(t *testing.T) {
		parkingArea := parkinglot.NewParkingArea(2)
		car := "UP-78-EB-1234"

		err := parkingArea.ParkIn(car)

		assert.Nil(t, err)
	})
	t.Run("should be able to un-park a car", func(t *testing.T) {
		parkingArea := parkinglot.NewParkingArea(2)
		car := "UP-78-EB-1234"
		_ = parkingArea.ParkIn(car)

		err := parkingArea.ParkOut(car)

		assert.Nil(t, err)
	})
	t.Run("should not be able to un-park a car when car is not parked", func(t *testing.T) {
		parkingArea := parkinglot.NewParkingArea(2)
		car := "UP-78-EB-1234"

		err := parkingArea.ParkOut(car)

		assert.NotNil(t, err)
	})
	t.Run("should not be able to park when all parking slots are full", func(t *testing.T) {
		parkingArea := parkinglot.NewParkingArea(2)
		car := "UP-78-EB-1234"
		otherCar := "UP-78-EB-1222"
		anotherCar := "UP-78-EB-1235"

		_ = parkingArea.ParkIn(car)
		_ = parkingArea.ParkIn(otherCar)

		err := parkingArea.ParkIn(anotherCar)

		assert.NotNil(t, err)

	})
}
