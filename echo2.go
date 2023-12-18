package main

import (
	"fmt"
	"os"
) 

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[4:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}