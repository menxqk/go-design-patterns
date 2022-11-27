package creational

import "testing"

func TestMotorkibeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type%d", sportBike.GetType())
}

func TestCarFactory(t *testing.T) {
	_, err := GetVehicleFactory(3)
	if err == nil {
		t.Fatal("Car factory with id 3 should not be recognized")
	}

	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.NewVehicle(LuxuryCarTpe)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats and %d wheels", carVehicle.GetSeats(), carVehicle.GetWheels())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury Car has %d doors.", luxuryCar.GetDoors())
}
