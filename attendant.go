package parkinglot

type Attendant struct {
	ParkingArea
}

func NewAttendant(parkingArea *ParkingArea) *Attendant {
	return &Attendant{*parkingArea}
}
