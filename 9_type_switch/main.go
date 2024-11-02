package main

import (
	"fmt"
)

func main() {
	var x interface{} = 2.3
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
}

// https://yourbasic.org/golang/find-type-of-object/
