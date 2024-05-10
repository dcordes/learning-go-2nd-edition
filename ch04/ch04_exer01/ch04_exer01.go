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
		if r == 100 {
			fmt.Println("match!")
		}
	}

	fmt.Println(random)
}
