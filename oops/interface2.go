package main

import (
    "fmt"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rect struct {
    Width  float64
    Height float64
}

func (r Rect) Area() float64 {
    return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %0.2f\n", s.Area())
    fmt.Printf("Perimeter: %0.2f\n", s.Perimeter())
}

func main() {
    r := Rect{
        Width:  5,
        Height: 10,
    }
    c := Circle{
        Radius: 7.5,
    }
    PrintShapeInfo(r)
    PrintShapeInfo(c)
}