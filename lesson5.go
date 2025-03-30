package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello from Goroutine!"
}

func main() {
	ch := make(chan string) // Create a channel
	go sendData(ch)         // Run goroutine
	message := <-ch         // Receive data
	fmt.Println(message)
}

// Hello from Goroutine!
// The output is Hello from Goroutine!.

// The main goroutine sends data to the channel. The sendData goroutine receives the data from the channel.

func main() {
	ch := make(chan int, 2) // Buffered channel with capacity 2

	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // ??
	fmt.Println(<-ch) // ??
}

// 1
// 2

// The output
