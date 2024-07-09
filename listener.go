package parkinglot

import "github.com/stretchr/testify/mock"

type Listener interface {
	fullSlots()
	spaceAvailable()
}

func (p *ParkingArea) notifyFullSlots() {
	for _, listener := range p.listener {
		listener.fullSlots()
	}
}

func (p *ParkingArea) notifySpaceAvailable() {
	for _, listener := range p.listener {
		listener.spaceAvailable()
	}
}

type TestHandRolledMockListener struct {
	lotFullNotificationReceived bool
}

func (t *TestHandRolledMockListener) LotFull() {
	t.lotFullNotificationReceived = true
}

type TestTestifyMockListener struct {
	mock.Mock
}

func (t *TestTestifyMockListener) LotFull() {
	t.Called()
}
