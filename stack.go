package main

import (
	"fmt"
)

func push(v int, a []int) {
	a = append(a, v)
}

func pop(v int, a []int) {
	index := 0
	for i, va := range a {
		if v == va {
			index = i
		}
	}
	a = append((a)[:index], (a)[index+1:]...)

}

func main() {
	a := [10]int{2, 3, 5, 1}
	push(67, a)
	fmt.Println(a)
	pop(3, a)
	fmt.Println(a)
	pop(1, a)
	push(88, a)
	push(100, a)
	fmt.Println(a)

}
