package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {

	res, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting response", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading response", err)
	// 	return
	// }

	// fmt.Println("Response : ", string(data))

	// convert json data to todo
	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error decoding response", err)
		return
	}

	fmt.Println("Response : ", todo)
	fmt.Println("Title Res : ", todo.Title)
	fmt.Println("Completed Res : ", todo.Completed)
}

func postRequest() {

	todo := Todo{
		UserId:    23,
		Title:     "Himanshu",
		Completed: true,
	}

	// convert todo to json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert data to reader
	jsonReader := bytes.NewReader([]byte(jsonString))

	res, err := http.Post(
		"http://jsonplaceholder.typicode.com/todos",
		"application/json",
		jsonReader,
	)

	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error getting response", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading response", err)
	// 	return
	// }

	// fmt.Println("Response : ", string(data))

	fmt.Println("Response status : ", res.Status)
}

func updateRequest() {

	todo := Todo{
		UserId:    23,
		Title:     "Himanshu Edited",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := bytes.NewReader([]byte(jsonString))

	res, err := http.NewRequest(
		http.MethodPut,
		"http://jsonplaceholder.typicode.com/todos/1",
		jsonReader,
	)

	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}

	res.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(res)

	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error sending request", resp.Status)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response : ", string(data))
	fmt.Println("Response status : ", resp.Status)
}

func deleteRequest() {

	res, err:= http.NewRequest(
		http.MethodDelete, 
		"http://jsonplaceholder.typicode.com/todos/1", 
		nil,
	)

	if err != nil {
		fmt.Println("Error deleting request", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(res)

	if err != nil {
		fmt.Println("Error deleting request", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error deleting request", resp.Status)
		return
	}

	fmt.Println("Response status : ", resp.Status)
}

func main() {

	// getRequest()
	// postRequest()
	// updateRequest()
	deleteRequest()
}
