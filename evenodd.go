package main

import (
	"fmt"
)

func evenN(even, odd, donech chan bool) {
	for i := 0; i <= 10; i += 2 {
		<-even // wait for the signal
		fmt.Println(i, "is even")
		if i == 10 {
			close(even)
			close(odd)
			donech <- true
			break
		}
		odd <- true
	}
}

func oddN(even, odd chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-odd // wait for the signal
		fmt.Println(i, "is odd")
		even <- true
	}
}

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	donech := make(chan bool)
	go evenN(even, odd, donech)
	go oddN(even, odd)
	even <- true
	<-donech // wait for the done signal
}
