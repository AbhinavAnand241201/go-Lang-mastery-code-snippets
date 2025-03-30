package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers()           // Start a goroutine
	time.Sleep(3 * time.Second) // Wait to let goroutine finish
	fmt.Println("Main function ends")
}

// 1
// 2
// 3
// 4
// 5
// Main function ends
