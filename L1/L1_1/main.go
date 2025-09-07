package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name string
	Age  int
}

func NewHuman(name string, age int) Human {
	return Human{Name: name, Age: age}
}

func (h *Human) SayHello() {
	fmt.Printf("Hello! My name is %s. I am %d years old.\n", h.Name, h.Age)

}

type Action struct {
	Human
}

func (a Action) ScreamHello() {
	fmt.Printf("HELLO!!! MY NAME IS %s!!! AND I AM %d YEARS OLD!!!", strings.ToUpper(a.Name), a.Age)
}

func main() {
	h1 := NewHuman("Vasya", 23)
	h2 := NewHuman("Sanya", 55)
	a1 := Action{Human: h2}

	h1.SayHello() // Vasya
	a1.SayHello() // Sanya

	a1.ScreamHello()

}
