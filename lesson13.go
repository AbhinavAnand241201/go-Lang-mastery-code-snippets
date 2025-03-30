// Flags Used:
// os.O_APPEND → Append new data to the file without erasing previous content.
// os.O_CREATE → Create the file if it does not exist.
// os.O_WRONLY → Open the file in write-only mode.

package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("fuzzing_errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		// log.Fatal() prints the error message and calls os.Exit(1) to stop execution.
	}
	defer file.Close()

	logger := log.New(file, "ERROR: ", log.LstdFlags)
	logger.Println("Crash detected at memory address 0xdeadbeef")
	logger.Println("Test case: input='AAAA'").    
}  

// defer file.close() is to close the file after the main funciton is done executing .




// 🔍 Expected Log File Output (fuzzing_errors.log)

// ERROR: 2025/03/29 15:30:12 Crash detected at memory address 0xdeadbeef
// ERROR: 2025/03/29 15:30:12 Test case: input='AAAA'
// Each line has:
// "ERROR: " prefix.
// A timestamp (provided by log.LstdFlags).
// The error message.