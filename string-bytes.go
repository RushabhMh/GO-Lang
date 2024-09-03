package main

import (
	"fmt"
)

func main() {
	// Example string
	str := "Hello, World!"
	// Convert string to byte slice
	byteSlice := []byte(str)
	// if we do it with interfaces we can do it with any type

	// Print the byte slice
	fmt.Println(byteSlice)

	// Print the byte slice as a string
	fmt.Println(string(byteSlice))
}
