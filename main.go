package main

import (
	"fmt"
)

func Calculate(a int) (result int) {
	result = a + 2
	return result
}

func main() {
	fmt.Println("Starting with Testing")
	result := Calculate(2)
	fmt.Println(result)
}
