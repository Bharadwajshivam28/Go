package main 

import "fmt"

const Pi = 3.14

func main () {
	const World= "Coder"
	fmt.Println("Hello", World)
	fmt.Println("Hi", Pi)

	const Truth = true
    fmt.Println("go rules", Truth)
}

func anotherFunction(){
	const World = "Sher"
	fmt.Println("Another message is", World)
}