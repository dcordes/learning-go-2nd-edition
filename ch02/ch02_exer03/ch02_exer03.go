package main

import "fmt"

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2_147_483_647
	bigI = 18_446_744_073_709_551_615

	b++
	smallI++
	bigI++

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
