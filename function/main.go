package main

import (
	"fmt"
)

func simpleFunc() {
	fmt.Println("Simple function")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Functions in GoLang")
	simpleFunc()
	fmt.Println(add(1, 2))
	fmt.Println(multiply(1, 2))
}
