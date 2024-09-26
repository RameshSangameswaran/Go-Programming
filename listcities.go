package main

import (
	"fmt"
	"os"
)

func main() {
	// Define list of new cities
	newCities := []string{
		"London",
		"Paris",
		"Rome",
		"Berlin",
		"Madrid",
	}

	// Create a new file to store the list of new cities
	file, err := os.Create("new_cities.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the list of new cities to the file
	for _, city := range newCities {
		fmt.Fprintf(file, "%s\n", city)
	}

	// Print a message to indicate that the list of new cities has been saved
	fmt.Println("List of new cities saved to new_cities.txt")
}
