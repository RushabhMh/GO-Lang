// The generator pattern is a design pattern that allows you to generate a
// sequence of values on demand. In Go, you can implement this pattern using
// channels and goroutines. The generator function will run in a separate goroutine
//  and send values to a channel, which can then be consumed by the caller.

// Example: Generator Pattern in Go
// Let's implement a simple generator that generates a sequence of integers.

package main

import (
	"fmt"
)

// IntGenerator generates a sequence of integers and sends them to a channel
func IntGenerator(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	// Create a generator that generates integers from 1 to 10
	gen := IntGenerator(1, 10)

	// Consume the generated values
	for value := range gen {
		fmt.Println(value)
	}
}
