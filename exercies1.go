package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func increment(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as completed when function exits

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Random delay
	// rand.Intn(n) returns a random number between 0 and n-1.
	// time.Duration converts the random number to a time.Duration value.
	// time.Sleep pauses the execution of the goroutine for the specified duration.

	mutex.Lock()
	counter++
	fmt.Printf("Goroutine %d incremented counter to: %d\n", id, counter)
	mutex.Unlock()
}

// imagine 10 waiters all running towards the cash register to change the cash flow ..but only gets to change the cash flow at one time .....so mutex.lock lock locks the register and only one waiter can change the cash flow that is the current go routine only...ok ?/?
func main() {
	rand.Seed(time.Now().UnixNano())
	// time.Now().UnixNano() is used to generate a random seed value for the random number generator.
	// unixnano give the current time in nanoseconds
	// Seed the random number generator with the current time in nanoseconds.

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go increment(i, &wg)
		//  this increment is a go routine

		// go increment(i) // this is a normal function call
		//  it's not a goroutine, it runs in the main goroutine and does not block the main goroutine.
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// the defer statement is used to mark a function call to be run when the surrounding function exits.

// The defer statement is often used to release resources acquired by a function, such as closing a file or unlocking a mutex.

// Goroutine 3 incremented counter to: 1
// Goroutine 1 incremented counter to: 2
// Goroutine 5 incremented counter to: 3
// Goroutine 2 incremented counter to: 4
// Goroutine 4 incremented counter to: 5
// Final Counter: 5

// // also this is also valid
// Goroutine 2 incremented counter to: 1
// Goroutine 4 incremented counter to: 2
// Goroutine 1 incremented counter to: 3
// Goroutine 5 incremented counter to: 4
// Goroutine 3 incremented counter to: 5
// Final Counter: 5
