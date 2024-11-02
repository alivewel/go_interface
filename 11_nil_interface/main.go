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

	var d *Dog
	fmt.Println(d == nil) // true

	i = d
	// type == *Dog, value == nil
	// поскольку type != nil, то i != nil
	fmt.Println(i == nil) // false
}
