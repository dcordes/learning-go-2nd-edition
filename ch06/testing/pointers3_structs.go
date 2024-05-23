package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// ~~~~~~~~~~~~~~~~~~

// note how it works if the struct is NOT dereferenced
func struct_test_variant1(vir *Person) {
	vir.age++
	//(*vir).age++
}

// but it also works if the struct IS dereferenced
func struct_test_variant2(vir *Person) {
	//vir.age++
	(*vir).age++
}

// why the equivalency?  because "Selectors automatically dereference pointers to structs"
//	https://go.dev/tour/moretypes/4
//	https://stackoverflow.com/questions/13533681/when-do-gos-pointers-dereference-themselves

// ~~~~~~~~~~~~~~~~~~

func int_test(virnum *int) {
	//virnum++	// this doesn't work -- no need to derefernce for an int
	*virnum++
}

func main() {
	fmt.Println("Do we need to deference a pointer to a struct in a function?")

	fred := Person{
		name: "Freddy",
		age:  10,
	}

	fmt.Println(fred.age)

	struct_test_variant1(&fred)
	fmt.Println(fred.age)
	struct_test_variant2(&fred)
	fmt.Println(fred.age)

	mynum := 17
	fmt.Println(mynum)
	int_test(&mynum)
	fmt.Println(mynum)

}
