package main

import (
	"fmt"
	"os"
)

func main() {

	// create file
	// file, err := os.Create("Example.txt")

	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }
	// defer file.Close()

	// fmt.Println("Successfully created file")

	// write to file
	// content := "Hello Himanshu Maurya"
	// byte, error := io.WriteString(file, content+"\n")

	// fmt.Println("Byte Written : ", byte)

	// if error != nil {
	// 	fmt.Println("Error while writing file", err)
	// 	return
	// }

	// fmt.Println("Successfully written to file")

	// open file
	// file, err := os.Open("Example.txt")

	// if err != nil {
	// 	fmt.Println("Error while opening file", err)
	// 	return
	// }
	// defer file.Close()

	// fmt.Println("Successfully opened file")

	// create a buffer to read content
	// buffer := make([]byte, 1024)

	// // read file
	// for {
	// 	n, err := file.Read(buffer)

	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		fmt.Println("Error while reading file", err)
	// 		return
	// 	}

	// 	fmt.Print(string(buffer[:n]))
	// }

	// read the entire file into byte slice
	content, err := os.ReadFile("Example.txt")

	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}

	fmt.Print(string(content))
}
