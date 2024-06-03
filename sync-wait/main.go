package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", i)
	// simulate some work
	fmt.Printf("Worker %d done\n", i)
}


func main() {

	var wg sync.WaitGroup

	// 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment the weightgroup counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Workers task completed")
}
