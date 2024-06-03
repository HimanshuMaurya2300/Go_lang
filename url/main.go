package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("Learning URL")

	myURL := "https://www.google.com/search?q=Himanshu+Maurya"
	fmt.Printf("Type of url is %T\n", myURL)

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}

	fmt.Println("Parsed URL is : ", parsedURL)
	fmt.Println("Scheme : ", parsedURL.Scheme)
	fmt.Println("Host : ", parsedURL.Host)
	fmt.Println("Path : ", parsedURL.Path)
	fmt.Println("RawQuery : ", parsedURL.Query())

	// modify url component
	parsedURL.Path = "/search"
	parsedURL.RawQuery = "q=Himanshu"

	newUrl := parsedURL.String()
	fmt.Println("New URL is : ", newUrl)
}
