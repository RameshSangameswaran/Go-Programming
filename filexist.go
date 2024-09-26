package main

import (
	"fmt"
	"os"
)

func main() {
	// Replace "yourfile.txt" with the path to the file you want to check
	filePath := "/home/nmit/a.py"

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", filePath)
	} else if err != nil {
		fmt.Printf("Error checking file: %s\n", err)
	} else {
		fmt.Printf("File exists: %s\n", filePath)
	}
}
