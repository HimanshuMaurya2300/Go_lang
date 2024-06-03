package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("JSON in GoLang")

	var person Person

	person.Name = "Himanshu"
	person.Age = 24
	person.IsAdult = true

	fmt.Println("Person data is : ", person)

	// convert person to json
	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error mashalling : ", err)
		return
	}

	fmt.Println("Json data is : ", string(jsonData))


	// unmarhalling
	var peronData Person

	err = json.Unmarshal(jsonData, &peronData)

	if err != nil {
		fmt.Println("Error unmarshalling : ", err)
		return
	}

	fmt.Println("Person data after unmarshalling is : ", peronData)
}
