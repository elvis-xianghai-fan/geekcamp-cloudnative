package main

import (
	"fmt"
	"reflect"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	// var ms struct1
	ms := &struct1{10, 15.5, "Chris"}
	fmt.Println(ms)

	arr1 := [3]int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println(reflect.TypeOf(arr1).String())
	fmt.Println(reflect.TypeOf(slice2).String())
}
