package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Println("read number", i, "from stdin")
	if i <= 10 {
		fmt.Println("The number", i, "is between 1 and 10")
	}
}
