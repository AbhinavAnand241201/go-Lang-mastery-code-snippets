package main

import (
	"fmt"
	"time"
)

func process(id int, ch chan string) {
	time.Sleep(time.Duration(id) * time.Second)
	// id into 1 billion nanoseconds sleep time
	// one main go routine and three parallel seperate go routines running in background
	ch <- fmt.Sprintf("Process %d done", id)
}

func main() {
	ch := make(chan string, 3)
	//  a capacity of 3 ....

	for i := 1; i <= 3; i++ {
		go process(i, ch)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
