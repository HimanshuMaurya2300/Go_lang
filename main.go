package main

import (
	"fmt"
	"mylearning/myUtil"
)

func main() {

	fmt.Println("Hello, World!")
	myUtil.Util()

	// Variables
	var age int = 24
	fmt.Println(age)

	var name string = "Himanshu"
	fmt.Println(name)

	var version = "1.0.0"
	fmt.Println(version)

	var dimension float64 = 56.78
	fmt.Println(dimension)

	var active bool = true
	active = false
	fmt.Println(active)

	const pi = 3.14
	fmt.Println(pi)

	person := 123
	person = 456
	fmt.Println(person)

	var Public = "public"
	var Private = "private"
	fmt.Println(Public, Private)
}
