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

	for i := 0; i < 100; i += 2 {
		<-even
		fmt.Println("Even :- ", i)
		time.Sleep(500 * time.Millisecond)
		odd <- true

	}

	select {

	case <-ctx.Done():
		fmt.Println("Stopping channel")
		return
	}
}

func printOdd(ctx context.Context, even, odd chan bool) {

	for i := 1; i < 100; i += 2 {
		<-odd
		fmt.Println("Odd:- ", i)
		time.Sleep(500 * time.Millisecond)
		even <- true

	}

	select {

	case <-ctx.Done():
		fmt.Println("Stopping channel")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	even := make(chan bool)
	odd := make(chan bool)

	go printEven(ctx, even, odd)
	go printOdd(ctx, even, odd)
	even <- true

	time.Sleep(2 * time.Second)

	cancel()

}
