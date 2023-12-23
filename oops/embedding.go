package main

import "fmt"

type Vehicle struct {
	Wheels int
}

func (v *Vehicle) Drive() {
	fmt.Println("Driving")
}

type Car struct {
	Vehicle
	Doors int
}

func (c *Car) Drive() {
	fmt.Println("Driving on a road")
}

func NewCar(doors int) *Car {
	c := Car{
		Vehicle{Wheels: 4},
		Doors: doors,
	}
	return &c
}

func main() {
	c := NewCar(4)
	fmt.Println("A car has", c.Wheels, "wheels")
	c.Drive()
}
