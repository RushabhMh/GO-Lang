package main

import "fmt"

type queue[T any] struct {
	items []T
}

func (q *queue[T]) push(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) pop() T {
	if len(q.items) == 0 {
		var t T
		return t
	}
	popitem := q.items[0]
	q.items = q.items[1:]
	return popitem
}

func main() {

	arr := &queue[int]{}
	arr.push(10)
	fmt.Println(arr.items)
	arr.push(20)
	fmt.Println(arr.items)
	arr.push(30)
	fmt.Println(arr.items)
	arr.pop()
	fmt.Println(arr.items)

}
