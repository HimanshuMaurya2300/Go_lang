package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Request in Golang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of res is %T\n", res)
	// fmt.Println(res)

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}

	fmt.Println(data)
	fmt.Println("Response : ", string(data))
}
