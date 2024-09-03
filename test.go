package main

import "fmt"

func checkEvenOrOdd(evenCh, oddCh chan int) {
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			evenCh <- n
		} else {
			oddCh <- n
		}
	}
	close(evenCh)
	close(oddCh)
	// tch <- 0
}

func main() {
	evenCh, oddCh := make(chan int), make(chan int)

	go checkEvenOrOdd(evenCh, oddCh)
	a, b := false, false
	for !a || !b {

		select {
		case v, ok := <-evenCh:
			if !ok {
				fmt.Println("Even channel closed")
				a = true
				evenCh = nil
				continue
			}
			fmt.Println(v, "is even")
		case v, ok := <-oddCh:
			if !ok {
				fmt.Println("Odd channel closed")
				b = true
				oddCh = nil
				continue
			}
			fmt.Println(v, "is odd")
			// case <-tch:
			//  break

		}
	}
}
