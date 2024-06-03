package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0

	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}

	numbers := []int{16, 25, 43, 48, 56}

	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, and value is %d\n", index, value)
	}

	data := "Himanshu"
	
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, char)
	}
}
