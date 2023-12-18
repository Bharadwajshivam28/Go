package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Index: %d, Value: %s\n", i, arg)
	}
}