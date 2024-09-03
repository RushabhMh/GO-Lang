package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generateprime(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
}

func printprime(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Prime numbers are:")
	ch := make(chan int)
	go generateprime(ch)
	printprime(ch)

}
