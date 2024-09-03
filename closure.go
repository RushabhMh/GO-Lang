package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}
}

func main() {
	fibonacci := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci())
	}

}
