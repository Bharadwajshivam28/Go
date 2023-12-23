package main

import "fmt"

// struct
type Person struct {
	Name string
	Age int
}
p := new(Person)
p.Name = "Shivam"
p.Age = 18 */

//struct with method Greet
type Person struct {
	name string
	age  int
}

func (p Person) Greet() string {
	return "Hello, my name is" + p.name
}

func main() {
	p := Person{name: "Shivam", age: 10}
	fmt.Println(p.Greet())
}

//
type Student struct {
	Name string
	roll int
}

func NewStudent(name string, age int) *Student{
	p := new(Student)
	p.Name = name
	p.Roll = p.roll
	return p
}