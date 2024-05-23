package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x *int

	fmt.Println(x == nil)
	fmt.Println(x)
	y := 10
	x = &y
	z := &y
	fmt.Println(*x)
	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println(&z)

	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(x))

	var a = new(int)
	var b = new(string)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
