package main

import "fmt"

func main() {
	fmt.Printf("type assertion\n~~~~~~~~~~\n")
	type abc any
	var first_var abc = 10
	var secon_var abc = "test"
	// imagine that the assignment of first and second var above took place elsewhere, and we want to perform an operation on them, but only if the types are right
	my1, ok := first_var.(int)
	if !ok {
		fmt.Println("incompatible type for first var")
	} else {
		fmt.Println(my1 + 2)
	}
	my2, ok := secon_var.(int)
	if !ok {
		fmt.Println("incompatible type for first var")
	} else {
		fmt.Println(my2 + 2)
	}

	fmt.Printf("type switch\n~~~~~~~~~~\n")

	my_type_switch := func(i any) {
		switch j := i.(type) {
		case nil:
			//
		case int:
			fmt.Println(j)
			fmt.Println(i.(int) + 9)
		case string:
			fmt.Println(j)
			fmt.Println(i.(string) + " hello")
		default:
			fmt.Println(j)
		}
	}

	my_type_switch(first_var)
	my_type_switch(secon_var)
}
