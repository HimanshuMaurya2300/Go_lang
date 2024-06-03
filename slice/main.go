package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = append(numbers, 11, 12, 13, 14, 15)
	fmt.Println("Numbers : ", numbers)
	fmt.Println("Length : ", len(numbers))
	fmt.Println("Capacity : ", cap(numbers))
	fmt.Printf("Data Type : %T\n", numbers)

	name := []string{}
	fmt.Println(name)

	nums := make([]int, 3, 4)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Println("Slice: ", nums)
	fmt.Println("Length : ", len(nums))
	fmt.Println("Capacity : ", cap(nums))
}
