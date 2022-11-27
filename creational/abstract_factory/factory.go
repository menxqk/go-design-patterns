package creational

import (
	"fmt"
)

// The Abstract Factory is level up to the factory method design pattern. It defines a
// factory of factories, adding a second level of abastraction for object creation.
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}

}

type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

type Car interface {
	GetDoors() int
}

type Motorbike interface {
	GetType() int
}

const (
	LuxuryCarTpe  = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (cf *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarTpe:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}
func (l *LuxuryCar) GetWheels() int {
	return 4
}
func (l *LuxuryCar) GetSeats() int {
	return 5
}

type FamilyCar struct{}

func (f *FamilyCar) GetDoors() int {
	return 5
}
func (f *FamilyCar) GetWheels() int {
	return 4
}
func (f *FamilyCar) GetSeats() int {
	return 5
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 1
}

func (c *CruiseMotorbike) GetType() int {
	return SportMotorbikeType
}
