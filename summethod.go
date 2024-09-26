package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(a, b float64) float64 {
	return a + b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the first number:")
	firstInput, _ := reader.ReadString('\n')
	firstNum, err := strconv.ParseFloat(firstInput[:len(firstInput)-1], 64) // Remove newline and convert
	if err != nil {
		fmt.Println("Invalid input for the first number.")
		return
	}

	fmt.Println("Enter the second number:")
	secondInput, _ := reader.ReadString('\n')
	secondNum, err := strconv.ParseFloat(secondInput[:len(secondInput)-1], 64) // Remove newline and convert
	if err != nil {
		fmt.Println("Invalid input for the second number.")
		return
	}

	result := sum(firstNum, secondNum)
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", firstNum, secondNum, result)
}
