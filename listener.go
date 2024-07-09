package parkinglot

type Listener interface {
	FullSlots()
	SpaceAvailable()
}

func (p *ParkingArea) notifyFullSlots() {
	for _, listener := range p.listener {
		listener.FullSlots()
	}
}

func (p *ParkingArea) notifySpaceAvailable() {
	for _, listener := range p.listener {
		listener.SpaceAvailable()
	}
}
