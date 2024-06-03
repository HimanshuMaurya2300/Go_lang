package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 43
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	// num = num + 1.23

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is ", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number := "1234g"
	num_int, err := strconv.Atoi(number)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("num_int is : ", num_int)
		fmt.Printf("Type of num_int is %T\n", num_int)
	}

	num_string := "3.14"
	num_float, err := strconv.ParseFloat(num_string, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("num_float is : ", num_float)
		fmt.Printf("Type of num_float is %T\n", num_float)
	}
}
