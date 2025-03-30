package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		// this is like putting a worker to sleep or giving him some rest ...
		// this is not a normal for-loop
		results <- job * job
	}
}

// Step-by-Step Explanation
// jobs is a channel that sends tasks (or values).
// for job := range jobs waits for values to arrive on jobs.
// When a value is sent to jobs, job receives it, and the loop body executes.
// When jobs is closed, the loop automatically exits.

func main() {
	const numWorkers = 3
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the workers
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close job channel to signal workers

	wg.Wait()
	close(results)

	// Read results
	for res := range results {
		fmt.Println("Result:", res)
		// We should never read from a closed channel, it will cause a panic.
		// so error
	}
}

// Worker 1 processing job 1
// Worker 2 processing job 2
// Worker 3 processing job 3
// Worker 1 processing job 4
// Worker 2 processing job 5
// Result: 1
// Result: 4
// Result: 9
// Result: 16
// Result: 25

// order may vary

// this is such an efficient concurrency
