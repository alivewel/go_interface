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
	var animal Animal
	fmt.Printf("Value %v, type %T\n", animal, animal) // Выведем значение и тип данной структуры
	if animal != nil {
		fmt.Println("animal is not nil")
	}

	// animal.Speak() // паника при попытке вызвать (interface == nil)
	// animal = "asf" // не скомпилируется, string does not implement Animal (missing method Speak)

	// var dog *Dog // если создаем таким образом, то ловим сегу при попытке заполнить поле Name
	dog := &Dog{}

	animal = dog
	animal.Speak() // OK
	fmt.Printf("Value %v, type %T\n", animal, animal)
	if animal != nil {
		fmt.Println("animal is not nil")
	}
	dog.Name = "Шайтан"
	fmt.Printf("Value %v, type %T\n", animal, animal)
}
