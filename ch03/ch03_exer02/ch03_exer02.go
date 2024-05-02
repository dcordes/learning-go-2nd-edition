package main

import (
	"fmt"
)

func main() {
	fmt.Println("chapter 03 -- exercise 02")

	message := "Hi 👩 and 👨"
	fmt.Println(message)
	//var b_message []byte = []byte(message)
	var r_message []rune = []rune(message)
	//fmt.Println(b_message)
	fmt.Printf("Here is the message as a rune array: %d\n", r_message)
	fmt.Printf("Fourth rune of that message, printed as a string: %s\n", string(r_message[3]))
}
