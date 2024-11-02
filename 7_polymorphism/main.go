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

// Функция, принимающая интерфейс Animal и вызывающая метод Speak
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{Name: "Бобик"}
	cat := Cat{Name: "Мурка"}

	// Вызов функции MakeAnimalSpeak для каждого животного
	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}