package main

import (
	"fmt"
)

func main() {

	// notice how when working with multi-dimensional arrays
	// that the type does NOT need to be specified when
	// specifying the values of the members of the 2nd array

	fmt.Println("hello world")

	mystr := []string{
		"aaa",
		"bbb",
		"ccc",
	}

	fmt.Println(mystr)

	for i, v := range mystr {
		fmt.Println(i, v)
	}

	mystrd := [][]string{
		{"aaa", "aaa1"},
		{"bbb", "bbb2"},
	}

	fmt.Println(mystrd)
	for i, v := range mystrd {
		fmt.Println(i, v)
	}
	for i, v := range mystrd {
		for i1, v1 := range v {
			fmt.Println(i, i1, v1)
		}
	}

	mystrdd := [][]string{
		[]string{"xxx", "xxx1"},
		[]string{"yyy", "yyy1"},
	}
	fmt.Println(mystrdd)
	for outer_i, outer_v := range mystrdd {
		for inner_i, inner_v := range outer_v {
			fmt.Println(outer_i, inner_i, inner_v)
		}
	}

	myint := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(myint)
	for outer_i, outer_v := range myint {
		for inner_i, inner_v := range outer_v {
			fmt.Println(outer_i, inner_i, inner_v)
		}
	}

	myint1 := [][]int{
		{7, 8, 9},
		{2, 1, 0},
	}
	fmt.Println(myint1)
	for outer_i, outer_v := range myint1 {
		for inner_i, inner_v := range outer_v {
			fmt.Println(outer_i, inner_i, inner_v)
		}
	}
}
