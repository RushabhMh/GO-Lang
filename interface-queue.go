package main

import "fmt"

type queue struct {
	items []interface{}
}

func (q *queue) push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *queue) pop() interface{} {
	if len(q.items) == 0 {
		return -1
	}
	popitem := q.items[0]
	q.items = q.items[1:]
	return popitem
}

func main() {

	arr := queue{items: []interface{}{1, 2, 3, 4}}
	arr.push(10)
	fmt.Println(arr.items)
	arr.push("rushabh")
	fmt.Println(arr.items)
	popitem := arr.pop()
	fmt.Println(popitem)
	fmt.Println(arr.items)
	popitem = arr.pop()
	fmt.Println(arr.items)
	arr.pop()
	fmt.Println(arr.items)
	arr.pop()
	fmt.Println(arr.items)
	arr.pop()
	fmt.Println(arr.items)

}
