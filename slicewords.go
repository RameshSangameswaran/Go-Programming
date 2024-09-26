package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a reader to read input from the console
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter two words separated by a space:")

	// Read the input
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any extra whitespace and split the input
	input = strings.TrimSpace(input)
	words := strings.Fields(input)

	// Check if we have exactly two words
	if len(words) != 2 {
		fmt.Println("Please enter exactly two words.")
		return
	}

	// Print the two words
	fmt.Printf("First word: %s\n", words[0])
	fmt.Printf("Second word: %s\n", words[1])
}
