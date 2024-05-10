package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := []int{}

	for i := 0; i < 100; i++ {
		r := rand.Intn(101)
		random = append(random, r)
	}

	for i := range random {
		message := "Never mind"

		if i != 0 {
			switch {
			case i%2 == 0 && i%3 == 0:
				message = "Six!"
			case i%2 == 0:
				message = "Two!"
			case i%3 == 0:
				message = "Three!"
			}
		}
		fmt.Println(i, message)
	}
}
