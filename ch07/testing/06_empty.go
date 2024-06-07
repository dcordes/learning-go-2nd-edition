package main

import "fmt"

func main() {
	var not_empty int

	not_empty = 5
	// the next line fails on a "go fmt" check
	//not_empty = "test"
	fmt.Println(not_empty)

	//
	var empty interface{}
	empty = 5
	empty = "test"
	fmt.Println(empty)

	// the newer method, using "any"
	var empty1 any
	empty1 = 51
	empty1 = "test1"
	fmt.Println(empty1)
}
