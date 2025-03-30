package main

import "fmt"

func main() {
	funcs := []func(){}

	for i := 0; i < 5; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}
}

// funcs is a slice of functions
// A slice is like Java’s ArrayList—a dynamic array that can grow.
// this function func(){} has no arguments and no return type
// so lenfuncs is 0 and capacity of func is 0

// 5 5 5 5 5
// The output is 5 5 5 5 5.
// The reason is that the anonymous function we append to funcs captures the variable i by reference, not by value.



for i := 0; i < 5; i++ {
    i := i  // This creates a new local copy of `i`
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}



// 0 1 2 3 4 is the output here because we are creating a new local copy of i in each iteration of the loop.