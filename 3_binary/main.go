package main

import (
	"fmt"
	"strconv"
)

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

func ToString(s fmt.Stringer) string {
	return s.String()
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	var num Binary = 200
	fmt.Printf("Value %v, type %T\n", num, num) // Выведем значение и тип данной структуры
	// Получим Value 11001000, type main.Binary)
	fmt.Println(ToString(num)) // Выводит: "11001000"
	fmt.Println(num)           // Выводит: "11001000"
}

// https://research.swtch.com/interfaces
