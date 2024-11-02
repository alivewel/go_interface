package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I

	i = &T{"Hello"}
	fmt.Printf("Value %v, type %T)\n", i, i)
	i.M()
}
