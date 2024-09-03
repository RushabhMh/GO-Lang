package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	ch <- 5

	for i := 2; i <= 10; i++ {
		go func() {
			ch <- i * 5
		}()
		fmt.Println(<-ch)
	}
	fmt.Println(<-ch)
	close(ch)

}
