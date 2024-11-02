package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue) // "[]int [1 2 3]"
}

// https://yourbasic.org/golang/find-type-of-object/
