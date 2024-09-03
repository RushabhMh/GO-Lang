package main

import (
	"fmt"
	
)

func printNumber(i int, ch chan int) {
	ch <- i
	// wg.Done()
}

func push(i int, ch chan int) {
	ch <- i

}

// func pop(ch chan int, wg *sync.WaitGroup) {

// 	fmt.Println(<-ch)
// 	runtime.Gosched()
// 	wg.Done()
// }

func main() {
	var n int
	var ch = make(chan int, 1)
	// var wg sync.WaitGroup
	fmt.Print("Enter number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}
	ch <- 1
	for i := 2; i <= n; i++ {
		// wg.Add(1)
		fmt.Println(<-ch)
		go push(i, ch)
	}

	// wg.Wait()
	fmt.Println()
}
