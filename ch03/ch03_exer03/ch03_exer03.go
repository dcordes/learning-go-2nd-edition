package main

import (
	"fmt"
)

func main() {
	fmt.Println("chapter 03 -- exercise 03")

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	abe := Employee{
		"abe",
		"lincoln",
		123,
	}
	bill := Employee{
		firstName: "bill",
		lastName:  "clinton",
		id:        456,
	}
	var jill Employee
	jill.firstName = "jill"
	jill.lastName = "jones"
	jill.id = 789

	fmt.Println(abe)
	fmt.Println(bill)
	fmt.Println(jill)
}
