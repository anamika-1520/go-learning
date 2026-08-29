package main

import (
	"fmt"
	"sync"

	"golang/goroutines"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(501)
	for i := 0; i <= 500; i++ {
		go goroutines.Worker(&wg)
		fmt.Printf(" Worker %v started\n", i)
	}
	wg.Wait()

	fmt.Println("Main finished")
}
