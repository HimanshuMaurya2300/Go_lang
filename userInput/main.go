package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Printf("Enter your name: ")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	myName, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", myName)
}
