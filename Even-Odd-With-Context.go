// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"time"
)

var counter int

func printEven(ctx context.Context, even, odd chan bool) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("Stopping even channel")
			time.Sleep(800 * time.Millisecond)
			return
		default:
			<-even
			fmt.Println("Even :- ", counter)
			counter++
			time.Sleep(800 * time.Millisecond)
			odd <- true
		}
	}

}

func printOdd(ctx context.Context, even, odd chan bool) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("Stopping odd channel")
			return
		default:
			<-odd
			fmt.Println("Odd:- ", counter)
			counter++
			even <- true
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	even := make(chan bool)
	odd := make(chan bool)

	go printEven(ctx, even, odd)
	go printOdd(ctx, even, odd)
	even <- true

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Main done")
}
