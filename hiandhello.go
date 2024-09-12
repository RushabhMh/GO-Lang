// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func PrintHi(Hi, Hello chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-Hi

		fmt.Println("Hi")
		Hello <- true
	}

}

func PrintHello(Hi, Hello chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-Hello
		if i == 9 {
			close(Hi)
			close(Hello)
			break
		}
		fmt.Println("Hello")
		Hi <- true
	}

}

func main() {
	Hi := make(chan bool)

	Hello := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)
	go PrintHi(Hi, Hello, &wg)
	go PrintHello(Hi, Hello, &wg)

	Hi <- true
	wg.Wait()
}
