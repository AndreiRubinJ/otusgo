package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	numWorkers := 10
	waitGroup.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go task(i, &waitGroup, &mutex, &counter)
	}
	waitGroup.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

func task(id int, wg *sync.WaitGroup, mu *sync.Mutex, counter *int) {
	defer wg.Done()
	mu.Lock()
	*counter++
	fmt.Printf("Worker %d incremented counter to %d\n", id, *counter)
	mu.Unlock()
}
