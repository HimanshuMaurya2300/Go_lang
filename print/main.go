package main

import "fmt"

func main() {

	age := 24
	name := "Himanshu"
	height := 5.8566
	fmt.Println("age:", age, "name:", name, "height:", height)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("name is %s\n", name)
	fmt.Printf("height is %.2f\n", height)

	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of height is %T\n", height)
}
