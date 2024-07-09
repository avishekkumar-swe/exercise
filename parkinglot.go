package parkinglot

import "errors"

type vehicleId string

type Vehicle struct {
	id vehicleId
}
type ParkingArea struct {
	totalSlots int
	listener   []Listener
	vehicles   []*Vehicle
}

func NewParkingArea(totalSlots int) *ParkingArea {
	return &ParkingArea{totalSlots, make([]Listener, 0), make([]*Vehicle, 0)}
}

func NewParkingAreaWithListener(totalSlots int, listener ...Listener) *ParkingArea {
	return &ParkingArea{totalSlots, listener, make([]*Vehicle, 0)}
}

func NewVehicle(id string) *Vehicle {
	return &Vehicle{vehicleId(id)}
}

func (p *ParkingArea) isAlreadyPresent(id string) bool {
	for _, thisVehicle := range p.vehicles {
		if thisVehicle.id == vehicleId(id) {
			return true
		}
	}
	return false
}
func (p *ParkingArea) ParkIn(id string) error {
	if p.isAlreadyPresent(id) {
		return errors.New("vehicle already parked")
	}
	if p.totalSlots == len(p.vehicles)+1 {
		p.notifyFullSlots()
	}
	if p.totalSlots > len(p.vehicles) {
		vehicle := NewVehicle(id)
		p.vehicles = append(p.vehicles, vehicle)
		return nil
	}
	return errors.New("parking area is full")
}
func (p *ParkingArea) ParkOut(id string) error {
	if p.isAlreadyPresent(id) == false {
		return errors.New("vehicle is not found")
	}
	if p.totalSlots == len(p.vehicles) {
		p.notifySpaceAvailable()
	}
	for index, vehicle := range p.vehicles {
		if vehicle.id == vehicleId(id) {
			p.vehicles = append(p.vehicles[:index], p.vehicles[index+1:]...)
			break
		}
	}
	return nil
}
