package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "iteration", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	go task("C")
	time.Sleep(4 * time.Second) // Give time for all goroutines to execute
}

// A iteration 1
// B iteration 1
// C iteration 1
// A iteration 2
// B iteration 2
// C iteration 2
// A iteration 3
// B iteration 3
// C iteration 3
