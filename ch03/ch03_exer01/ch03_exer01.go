package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	slice_one := greetings[:2]
	slice_two := greetings[1:4]
	slice_thr := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(slice_one)
	fmt.Println(slice_two)
	fmt.Println(slice_thr)
}
