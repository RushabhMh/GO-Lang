package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insertatBeginning(data int) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
}

func (l *LinkedList) insertatEnd(data int) {
	newNode := &Node{data: data}
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

func (l *LinkedList) Delete(data int) {

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

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ->\t")
		current = current.next
	}
	fmt.Print("nil")

}

func main() {

	l := &LinkedList{}

	l.insertatBeginning(12)
	l.insertatBeginning(13)
	l.insertatBeginning(14)
	l.insertatBeginning(15)
	l.Display()

}
