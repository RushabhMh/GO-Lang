package main

import (
	"fmt"
	"sync"
)

func FanIn(channels ...chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))

	output := func(in <-chan int) {
		for v := range in {
			out <- v
		}
		wg.Done()
	}

	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Done()
		close(out)
	}()

	return out
}
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 15; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	out := FanIn(ch1, ch2, ch3)

	for c := range out {
		fmt.Println(c)
	}

}
