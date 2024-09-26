package main

import "fmt"

func main() {
	var f float64
	fmt.Println("Enter a float:")
	fmt.Scanln(&f)

	i := int(f)
	fmt.Println("Truncated integer:", i)
}
