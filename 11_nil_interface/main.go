package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func main() {
	var i interface{}
	// type == nil, value == nil
	// поэтому i == nil
	fmt.Println(i == nil) // true
	fmt.Printf("Value %v, type %T\n\n", i, i)

	var d *Dog
	// type == *Dog, value == nil
	fmt.Println(d == nil) // true
	fmt.Printf("Value %v, type %T\n\n", d, d)

	i = d
	// type == *Dog, value == nil
	// поскольку type != nil, то i != nil
	fmt.Println(i == nil) // false
	fmt.Printf("Value %v, type %T\n", i, i)
}
