package parkinglot_test

import (
	parkinglot "github.com/avishekkumar-swe/exercise/tree/master"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestParkingLotOwnerListener(t *testing.T) {
	t.Parallel()
	t.Run("owner should be notify for full parking slots", func(t *testing.T) {
		owner := &TestHandRolledMockListener{}
		parkingArea := parkinglot.NewParkingAreaWithListener(1, owner)
		car := "UP-78-EB-1995"
		_ = parkingArea.ParkIn(car)

		assert.True(t, owner.lotFullNotificationReceived)
	})
	t.Run("owner should be notify for recent available parking slot after full parking slots", func(t *testing.T) {
		owner := &TestTestifyMockListener{}
		parkingArea := parkinglot.NewParkingAreaWithListener(1, owner)
		owner.On("FullSlots")
		owner.On("SpaceAvailable")
		car := "UP-78-EB-1995"
		_ = parkingArea.ParkIn(car)

		_ = parkingArea.ParkOut(car)

		owner.AssertCalled(t, "SpaceAvailable")
	})
}
func TestParkingLotListeners(t *testing.T) {
	t.Parallel()
	t.Run("owner and cop should be notify for full parking slots", func(t *testing.T) {
		owner := &TestHandRolledMockListener{}
		cop := &TestHandRolledMockListener{}
		parkingArea := parkinglot.NewParkingAreaWithListener(1, owner, cop)
		car := "UP-78-EB-1995"
		_ = parkingArea.ParkIn(car)

		assert.True(t, owner.lotFullNotificationReceived)
		assert.True(t, cop.lotFullNotificationReceived)
	})
	t.Run("owner and cop should be notify for recent available parking slot after full parking slots", func(t *testing.T) {
		owner := &TestTestifyMockListener{}
		cop := &TestTestifyMockListener{}
		parkingArea := parkinglot.NewParkingAreaWithListener(1, owner, cop)
		owner.On("FullSlots")
		owner.On("SpaceAvailable")
		cop.On("FullSlots")
		cop.On("SpaceAvailable")
		car := "UP-78-EB-1995"
		_ = parkingArea.ParkIn(car)

		_ = parkingArea.ParkOut(car)

		cop.AssertCalled(t, "SpaceAvailable")
		owner.AssertCalled(t, "SpaceAvailable")
	})
}

type TestHandRolledMockListener struct {
	lotFullNotificationReceived        bool
	spaceAvailableNotificationReceived bool
}

func (t *TestHandRolledMockListener) FullSlots() {
	t.lotFullNotificationReceived = true
}

func (t *TestHandRolledMockListener) SpaceAvailable() {
	t.spaceAvailableNotificationReceived = true
}

type TestTestifyMockListener struct {
	mock.Mock
}

func (t *TestTestifyMockListener) FullSlots() {
	t.Called()
}

func (t *TestTestifyMockListener) SpaceAvailable() {
	t.Called()
}
