package main

import (
	"fmt"
)

func main() {
	fmt.Println(true)
	true := false
	fmt.Println(true)

	fmt.Println(false)
	false := "aardvark"
	fmt.Println(false)

}
