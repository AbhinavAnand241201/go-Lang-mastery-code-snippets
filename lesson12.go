package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("fuzz_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("New test case executed...\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("New data appended successfully!")
}
