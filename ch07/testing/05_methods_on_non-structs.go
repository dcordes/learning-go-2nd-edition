package main

import "fmt"

type basicint int

func (b basicint) Triple() {
	fmt.Println(b * 3)
}

func main() {
	bi := basicint(6)

	fmt.Println(bi)

	bi.Triple()
}
