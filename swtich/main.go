package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	month := "August"

	switch month {
	case "December", "January", "February":
		fmt.Println("Winter")
	case "March", "April", "May", "June":
		fmt.Println("Summer")
	default:
		fmt.Println("Rainy")
	}

	temp := 45

	switch {
	case temp > 30:
		fmt.Println("Hot")
	case temp < 10:
		fmt.Println("Cold")
	default:
		fmt.Println("Normal")
	}
}
