package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("SayHello done")
}

func SayHi() {
	fmt.Println("Hi")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("SayHi done")
}

func main() {
	fmt.Println("Learning Goroutines")
	go sayHello()
	go SayHi()

	// wait for sayHello to finish
	time.Sleep(3000 * time.Millisecond)
}
