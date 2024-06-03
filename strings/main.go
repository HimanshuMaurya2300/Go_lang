package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Apple, Mango, Banana, Grapes"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three two five two"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "   Hello Himanshu !      "
	str = strings.TrimSpace(str)
	fmt.Println(str)

	str1 := "Himanshu"
	str2 := "Maurya"
	result := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("Result is : ", result)
}
