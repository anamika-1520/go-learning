package main

import (
	"fmt"
	"sync"
)

func square(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- n * n
}

func main() {

	var wg sync.WaitGroup

	ch := make(chan int)

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go square(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("Result:", result)
	}

	fmt.Println("All workers finished")
}
