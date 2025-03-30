package main

import (
	"fmt"
	"os"
)

func main() {
	content := "Fuzzing test case - Crash found!\n"

	err := os.WriteFile("fuzz_log.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Log written successfully!")
}

// permission 0644 allows read/write for owner and read for others ...





// package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("fuzz_log.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:\n", string(data))
}

// read data from a file 