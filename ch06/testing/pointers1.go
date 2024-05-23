package main

import "fmt"

func mod_val(myvar int) int {
	return myvar + 4
}

func mod_val1(myvar int) {
	myvar++
}

// notice here how no return is necessary
func mod_val2(myvar *int) {
	*myvar++
}

func mod_ptr(myvar *int) int {
	return *myvar + 5
}

func main() {
	a := 2
	a++
	fmt.Println((a))

	x := 10
	var y *int
	y = &x

	fmt.Println(x, y, *y)
	fmt.Println(mod_val(x))

	xx := 1
	fmt.Println("~~~", xx)
	mod_val1(xx)
	fmt.Println("~~~", xx)

	// notice in this block how the value of the soure variable
	//	changes, and that there is no need to return the value
	//	from the function
	yy := 1
	fmt.Println("===", yy)
	mod_val2(&yy)
	fmt.Println("===", yy)

	fmt.Println(mod_ptr(&x))

}
