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

func (a *Attendant) notifyFullSlots() {
	for _, listener := range a.listener {
		listener.FullSlots()
	}
}

func (a *Attendant) notifySpaceAvailable() {
	for _, listener := range a.listener {
		listener.SpaceAvailable()
	}
}
