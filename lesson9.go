// okay so lets do a detailed step by step dry run of this code spciifically // When multiple goroutines modify shared data, we must use a mutex to prevent data corruption.

package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// WaitGroup is a keyword ...
	// A WaitGroup waits for a collection of goroutines to finish.

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
		// pointer refernece ...so keep in mind ...

	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
