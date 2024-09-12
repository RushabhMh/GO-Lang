package main

import (
	"fmt"
)

// PrintNumber prints the number and signals the next goroutine.
func PrintNumber(n int, prevChan, nextChan chan struct{}) {
	<-prevChan       // Wait for signal from the previous goroutine
	fmt.Println(n)   // Print the number
	nextChan <- struct{}{} // Signal the next goroutine
}

func main() {
	n := 10 // Number of natural numbers to print

	// Create channels for synchronization between goroutines
	channels := make([]chan struct{}, n+1)
	for i := range channels {
		channels[i] = make(chan struct{})
	}

	// Launch goroutines
	for i := 1; i <= n; i++ {
		go PrintNumber(i, channels[i-1], channels[i])
	}

	// Start the sequence by signaling the first goroutine
	channels[0] <- struct{}{}

	// Wait for the last goroutine to finish
	<-channels[n]
	fmt.Println("All numbers printed in sequential order.")
}


// package main

// import (
// 	"fmt"
// )

// func main() {
// 	n := 10                    // Number of natural numbers to print
// 	nums := make(chan int, n)  // Channel to send numbers

// 	// Goroutine that prints numbers in order
// 	go func() {
// 		for num := range nums {
// 			fmt.Println(num)
// 		}
// 	}()

// 	// Send numbers to the channel in order
// 	for i := 1; i <= n; i++ {
// 		nums <- i
// 	}

// 	close(nums) // Close the channel after sending all numbers
// 	fmt.Println("All numbers printed in sequential order.")
// }


// mutex

package main

import (
	"fmt"
	"sync"
)

var (
	counter int         // Shared counter to keep track of the next number to print
	n       = 10        // Number of natural numbers to print
	mutex   sync.Mutex  // Mutex to control access to the counter
	wg      sync.WaitGroup // WaitGroup to wait for all goroutines to finish
)

func printSequentially(id int) {
	defer wg.Done() // Signal that this goroutine is done
	for {
		mutex.Lock() // Lock the critical section
		if counter >= n { // Exit condition
			mutex.Unlock() // Ensure the mutex is unlocked before exiting
			return
		}
		if counter == id { // Check if it's this goroutine's turn to print
			fmt.Println(counter + 1) // Print the number (1-based indexing)
			counter++                 // Increment the counter for the next turn
		}
		mutex.Unlock() // Unlock the critical section
	}
}

func main() {
	// Launch n goroutines
	for i := 0; i < n; i++ {
		wg.Add(1)
		go printSequentially(i)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All numbers printed in sequential order.")
}
