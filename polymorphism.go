package main

import "fmt"

type Speaker interface {
	Speak()
}

type Address struct {
	State string
	City  string
}

type Person struct {
	Name string
	Address
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is " + p.Name)
}

func Introduce(p Speaker) {
	p.Speak()
}

func main() {

	p := Person{
		Name: "Rushabh",
	}
	Introduce(p)

}
