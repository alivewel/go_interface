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

func (d Dog) Bark() string {
	return fmt.Sprintf("%s громко лает!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Кошка %s мяукает", c.Name)
}

func (c Cat) Purr() string {
	return fmt.Sprintf("%s мурлычет.", c.Name)
}

// функция в которой происходит type assertion
func processAnimalTypeAssertion(animal Animal) {
	if dog, ok := animal.(*Dog); ok {
		fmt.Printf("Type: %T Value: %#v\n", dog, dog)
		fmt.Println(dog.Bark())
	}
	if cat, ok := animal.(*Cat); ok {
		fmt.Printf("Type: %T Value: %#v\n", cat, cat)
		fmt.Println(cat.Purr())
	}
}

// функция в которой происходит type switch
func processAnimalTypeSwitch(animal Animal) {
	switch v := animal.(type) {
	case *Dog:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
		fmt.Println(v.Bark())
	case *Cat:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
		fmt.Println(v.Purr())
	default:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
	}
}

func main() {
	dog := &Dog{Name: "Бобик"}
	cat := &Cat{Name: "Мурка"}

	processAnimalTypeAssertion(dog)
	processAnimalTypeAssertion(cat)

	fmt.Println()

	processAnimalTypeSwitch(dog)
	processAnimalTypeSwitch(cat)
}
