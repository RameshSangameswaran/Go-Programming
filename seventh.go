package main

import "fmt"

func main() {
	var sum float32 = 0
	var i float32 = 1
	for i <= 10 {
		sum = sum + i
		i = i + 1
	}
	fmt.Println("sum of 1 to 10 numbers is ", sum)
}
