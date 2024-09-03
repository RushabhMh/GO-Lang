package main

import (
	"errors"
	"fmt"
)

// Stack is a stack data structure that holds items of type interface{}
type Stack struct {
	items []interface{}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	// Example usage with int
	intStack := &Stack{}
	intStack.Push(1)
	intStack.Push('a')
	intStack.Push(3)
	fmt.Println("Int Stack size:", intStack.Size())
	item, _ := intStack.Pop()
	fmt.Println("Popped item from Int Stack:", item)
	fmt.Println("Popped item from Int Stack:", intStack.items)

	// Example usage with string
	stringStack := &Stack{}
	stringStack.Push("a")
	stringStack.Push("b")
	stringStack.Push("c")
	fmt.Println("String Stack size:", stringStack.Size())
	itemStr, _ := stringStack.Pop()
	fmt.Println("Popped item from String Stack:", itemStr)
	fmt.Println("Popped item from String Stack:", stringStack.items)
}
