package Interface

func (c Car) Speed() uint {
	return c.MaxSpeed
}

func (c Car) Capacity() uint {
	return c.passenger
}

func (b Bike) Speed() uint {
	return b.MaxSpeed
}