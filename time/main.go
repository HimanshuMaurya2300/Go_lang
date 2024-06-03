package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println("Current Time : ", currTime)
	fmt.Printf("Type of currTime is %T\n", currTime)

	formatttedTime := currTime.Format("02-01-2006 15:04:05, Monday")
	fmt.Println("Formatted Time : ", formatttedTime)

	formatttedTime2 := currTime.Format("02-01-2006 3:04:05 PM")
	fmt.Println("Formatted Time : ", formatttedTime2)

	dateStr := "23/04/2024"
	formatted_time, _ := time.Parse("02/01/2006", dateStr)
	fmt.Println("Formatted Time : ", formatted_time)

	new_date := currTime.Add(24 * time.Hour)
	formatted_new_date := new_date.Format("02/01/2006 3:04:05 PM Monday")
	fmt.Println("New Date : ", formatted_new_date)
}
