package Interface

type Car struct {
	MaxSpeed uint
	passenger uint
}

func newCar(maxSpeed, passenger uint) *Car {
	return &Car{
		MaxSpeed: maxSpeed,
		passenger: passenger,
	}
}

type Bike struct {
	MaxSpeed uint
}

func newBike(maxSpeed uint) *Bike {
	return &Bike{
		MaxSpeed: maxSpeed,
	}
}