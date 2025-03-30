package main

import "fmt"

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Memory corruption detected!")
}

func main() {
	riskyFunction()
	fmt.Println("Program continues...")
}

// How panic and recover Work Together
// panic("Memory corruption detected!") is called.
// Stack unwinding starts.
// Before exiting riskyFunction, the deferred function executes.
// recover() catches the panic, preventing the program from crashing.
// "Recovered from panic: Memory corruption detected!" is printed.
// Execution resumes in main(), and "Program continues..." is printed.
