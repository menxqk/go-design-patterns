package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheel on a car must be 4 and they were %d", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d", car.Wheels)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d", motorbike.Wheels)
	}
	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s", motorbike.Structure)
	}
	if motorbike.Seats != 2 {
		t.Errorf("Seats on a motorbike must be 2 and was %d", motorbike.Seats)
	}

	busBuilder := &BusBuilder{}

	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("Wheels on a bus must be 8 and they were %d", bus.Wheels)
	}
	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus' and was %s", bus.Structure)
	}
	if bus.Seats != 45 {
		t.Errorf("Seats on a bus must be 45 and was %d", bus.Seats)
	}
}
