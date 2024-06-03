package main

import "fmt"

func modifyValueByReference(value *int) {
	*value = 20
}

func main() {
	var num int
	var ptr *int

	num = 10
	ptr = &num

	fmt.Println("Address of num : ", ptr)
	fmt.Println("Value of num : ", *ptr)

	nums := 2
	p := &nums
	fmt.Println(p)
	fmt.Println(*p)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println(value)
}
