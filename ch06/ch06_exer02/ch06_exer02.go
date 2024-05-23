package main

import "fmt"

func UpdateSlice(aStrSlice []string, aStr string) {
	fmt.Println("[UpdateSlice] 1: ", aStrSlice)
	aStrSlice[len(aStrSlice)-1] = aStr // YES -- it is UGLY to have to refer to the last element of the slice this way
	fmt.Println("[UpdateSlice] 2: ", aStrSlice)
}

func GrowSlice(aStrSlice []string, aStr string) {
	fmt.Println("[GrowSlice] 1: ", aStrSlice)
	aStrSlice = append(aStrSlice, aStr)
	fmt.Println("[GrowSlice] 2: ", aStrSlice)
}

func main() {
	mystrs1 := []string{"aaa", "bbb", "ccc"}
	fmt.Println("Before UpdateSlice: ", mystrs1)
	UpdateSlice(mystrs1, "zzz")
	fmt.Println("After UpdateSlice: ", mystrs1) // note that the additional value IS in the slice, because we changed an existing element in the slice rather than extending it -- see README

	fmt.Println()

	mystrs2 := []string{"aaa", "bbb", "ccc"}
	fmt.Println("Before GrowSlice: ", mystrs2)
	GrowSlice(mystrs2, "zzz")
	fmt.Println("After GrowSlice: ", mystrs2) // note that the additional value is NOT in the slice, because the LEN pointer in our original slice was not updated when the called function performed an APPEND -- see README
}
