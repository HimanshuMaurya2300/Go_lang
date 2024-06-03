package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of program")
	data := add(5, 6)
	defer fmt.Println("Data is : ", data)
	defer fmt.Println("Middle of program")
	fmt.Println("End of program")
}
