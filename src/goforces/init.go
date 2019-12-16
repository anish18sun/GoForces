package main

import (
	"fmt"
)

func start(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("The value:", start(4, 5))
}
