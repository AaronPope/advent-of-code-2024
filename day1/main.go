package main

import (
	"fmt"
	"os"
)

func TestPub() {
	fmt.Println("TESTING...")
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}

	file.Close()
}
