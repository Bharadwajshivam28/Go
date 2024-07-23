package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Valu:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Value:", m["Answer"])
 
	v, ok := m["Answer"]
	fmt.Println("Value:", v, "Present?", ok)

}