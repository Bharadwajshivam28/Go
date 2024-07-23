package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell": {40, 60},
	"Google": {10,10},
}

func main(){
	fmt.Println(m)
}