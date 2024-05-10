package main

import (
	"fmt"
)

func main() {
	samples := []string{"hello", "apple_🍎!"}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
	fmt.Println()
}
