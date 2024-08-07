package main

import(
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS")
    case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}