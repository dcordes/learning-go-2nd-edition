package main

import (
	"fmt"
)

func prefixer(prefix string) func(string) string {
	return func(myval string) string {
		return prefix + " " + myval
	}
}

func main() {
	fmt.Println("ch05_exer03")

	//

	new_func := prefixer("oscar")

	fmt.Println(new_func("jones"))

	//

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
