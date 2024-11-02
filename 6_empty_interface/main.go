package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Собака %s лает", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Кошка %s мяукает", c.Name)
}

func main() {
	dog := &Dog{Name: "Шайтан"}

	var emptyInterface interface{}

	fmt.Printf("Value %v, type %T\n", emptyInterface, emptyInterface)

	emptyInterface = dog

	fmt.Printf("Value %v, type %T\n", emptyInterface, emptyInterface)

	emptyInterface = 123

	fmt.Printf("Value %v, type %T\n", emptyInterface, emptyInterface)

	emptyInterface = true

	fmt.Printf("Value %v, type %T\n", emptyInterface, emptyInterface)
}
