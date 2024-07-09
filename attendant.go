package parkinglot

import "errors"

type ParkingLots map[*ParkingArea]bool

type Attendant struct {
	availableParkingArea ParkingLots
}

func NewAttendant(parkingArea ...*ParkingArea) *Attendant {
	attendant := &Attendant{make(ParkingLots)}
	attendant.addParkingLots(parkingArea...)
	return attendant
}

func (a *Attendant) addParkingLots(parkingArea ...*ParkingArea) {
	for _, area := range parkingArea {
		a.availableParkingArea[area] = true
	}
}

func (a *Attendant) ParkIn(carId string) error {
	for parkingArea, _ := range a.availableParkingArea {
		if parkingArea.ParkIn(carId) == nil {
			return nil
		}
	}
	return errors.New("parking area not found")
}

func (a *Attendant) ParkOut(carId string) error {
	for parkingArea, _ := range a.availableParkingArea {
		if parkingArea.ParkOut(carId) == nil {
			return nil
		}
	}
	return errors.New("car not found in any parking area")
}
