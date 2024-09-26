package main

import (
	"fmt"
)

// Method to calculate the square of a number
func square(num int) int {
	return num * num
}

// Method that calls the square method and adds a constant value
func calculate(num int, constant int) int {
	squaredValue := square(num) // Call the square method
	return squaredValue + constant
}

func main() {
	var number, constant int

	// Read input from the user
	fmt.Println("Enter a number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Invalid input for the number.")
		return
	}

	fmt.Println("Enter a constant value to add:")
	_, err = fmt.Scan(&constant)
	if err != nil {
		fmt.Println("Invalid input for the constant value.")
		return
	}

	// Call the calculate method
	result := calculate(number, constant)
	fmt.Printf("The result of squaring %d and adding %d is %d\n", number, constant, result)
}
