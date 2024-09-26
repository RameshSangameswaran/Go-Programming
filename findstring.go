package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I am a batman"
	if strings.HasPrefix(str, "I") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("String found", str)
	} else {
		fmt.Println("String not found", str)
	}
}
