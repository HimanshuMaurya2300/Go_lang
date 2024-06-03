package main

import "fmt"

func divide1(a, b float64) float64 {
	return a / b
}

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func divide3(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "cannot divide by zero"
	}

	return a / b, "nil"
}

func main() {
	fmt.Println("Error Handling")

	ans := divide1(10, 4)
	fmt.Println(ans)

	ans, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

	ans, error := divide3(10, 0)
	if(error != "nil"){
		fmt.Println(error)
	}else{
		fmt.Println(ans)
	}
}
