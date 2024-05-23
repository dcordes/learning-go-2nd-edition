package main

import (
	"fmt"
)

func main() {
	// you can get the pointer address of something that is not a primitive
	//	literal (numbers, booleans, or strings) or a constant -- like a slice
	x := []int{4, 5, 6}
	fmt.Println(x[1])

	y := &[]int{7, 8, 9}
	// note the format here for dereferencing and accessing a slice element:
	//	NOT --> *y[0]
	//  BUT --> (*y)[0]
	fmt.Println((*y)[0])

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// on the other hand, you can't get a pointer address of a primitive literal,
	//	as said above, like an int or string or boolean
	//
	// Why not?  because anything that is not a composite literal is treated in the
	//	the same way as a CONSTANT, in the sense (I guess...) that it cannot be modified
	// -- I guess when we modify a primitive literal we are really just creating something
	// brand new each time and discarding the old one -- so the address is always changing
	// unlike a composite literal where the address stays the same and so where it makes
	// sense to use a fixed address, like a permanent residence
	z := &5
	fmt.Println(*z)
}
