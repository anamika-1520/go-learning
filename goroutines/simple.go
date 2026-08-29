package goroutines

import (
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker finished")
}
