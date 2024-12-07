package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numWorkers := 20
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go task(i, &wg, &mu, &counter)
	}

	wg.Wait()

	assert.Equal(t, numWorkers, counter, "Counter should equal the number of workers")
}

func TestNoRaceCondition(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("No race condition", func(t *testing.T) {
			var counter int
			var wg sync.WaitGroup
			var mu sync.Mutex

			numWorkers := 10
			wg.Add(numWorkers)

			for j := 1; j <= numWorkers; j++ {
				go task(j, &wg, &mu, &counter)
			}

			wg.Wait()

			assert.Equal(t, numWorkers, counter, "Counter should equal the number of workers")
		})
	}
}
