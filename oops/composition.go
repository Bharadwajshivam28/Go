package main

import (
    "fmt"
)

type Animal struct {
    Name string
    Age  int
}

func (a *Animal) Speak() {
    fmt.Println("Hello, my name is", a.Name)
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    dog := &Dog{
        Animal: Animal{
            Name: "Fido",
            Age:  3,
        },
        Breed: "Labrador",
    }

    dog.Speak()
    fmt.Println("I am a", dog.Breed)
}