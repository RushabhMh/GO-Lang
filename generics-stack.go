package main

import (
	"errors"
	"fmt"
)




// Stack is a generic stack data structure
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func main() {
	// Example usage with int
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println("Int Stack size:", intStack.Size())
	item, _ := intStack.Pop()
	fmt.Println("Popped item from Int Stack:", item)
	fmt.Println(intStack.items)
	// Example usage with string
	stringStack := &Stack[string]{}
	stringStack.Push("a")
	stringStack.Push("b")
	stringStack.Push("c")
	fmt.Println("String Stack size:", stringStack.Size())
	itemStr, _ := stringStack.Pop()
	fmt.Println("Popped item from String Stack:", itemStr)
	fmt.Println(stringStack.items)
}
