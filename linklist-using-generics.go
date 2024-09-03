package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) insertatBeginning(data T) {
	newNode := &Node[T]{data: data, next: l.head}
	l.head = newNode
}

func (l *LinkedList[T]) insertatEnd(data T) {
	newNode := &Node[T]{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

}

func (l *LinkedList[T]) Delete(data T) {
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}

}

func (l *LinkedList[T]) Display() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ->\t")
		current = current.next
	}
	fmt.Print("nil")

}

func main() {

	l := &LinkedList[int]{}

	l.insertatBeginning(12)
	l.insertatBeginning(13)
	l.insertatBeginning(14)
	l.insertatBeginning(15)
	l.Display()

}
