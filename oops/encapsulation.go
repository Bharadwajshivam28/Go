package main

import "fmt"

// Define a struct representing a car
type Car struct {
    brand  string
    model  string
    price  float64
    engine Engine
}

// Define a struct representing an engine
type Engine struct {
    horsepower int
    fuelType   string
}

// Define a method for the Car struct to get the price
func (c Car) GetPrice() float64 {
    return c.price
}

// Define a method for the Car struct to start the engine
func (c *Car) StartEngine() {
    fmt.Println("Starting the engine of", c.brand, c.model)
    // Additional logic can be added here to start the engine
}

func main() {
    // Create a new Car object
    myCar := Car{
        brand: "Toyota",
        model: "Camry",
        price: 25000.0,
        engine: Engine{
            horsepower: 180,
            fuelType:   "Gasoline",
        },
    }

    // Accessing data through encapsulation
    fmt.Println("Car brand:", myCar.brand)
    fmt.Println("Car model:", myCar.model)
    fmt.Println("Car price:", myCar.GetPrice())

    // Calling a method through encapsulation
    myCar.StartEngine()
}
