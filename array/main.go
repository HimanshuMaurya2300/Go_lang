package main

import "fmt"

func main() {
	fmt.Println("Array in Golang")

	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println(arr)

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	fmt.Println("Length of the array is : ", len(nums))
	fmt.Println("Value of the 3rd index is : ", nums[2])

	var price [5]int
	fmt.Println(price)

	var test [5]bool
	fmt.Println(test)

	var s [5]string
	fmt.Println(s)
	s[0] = "Himanshu"
	s[3] = "Raj"
	s[4] = "Maurya"
	fmt.Printf("String array is %q\n", s)
}
