package main

import (
	"fmt"
)

func main() {
	mylist := []string{"abe", "whit", "robin"}

	fmt.Println(mylist[0])
	for k, v := range mylist {
		fmt.Println(k, v)
	}
	for _, v := range mylist {
		fmt.Println(v)
	}
	for k := range mylist {
		fmt.Println(k)
	}

	//

	mymap := map[string]string{
		"oscar": "fred",
		"jim":   "bo",
		"mary":  "cassat",
	}
	for k, v := range mymap {
		fmt.Println(k, v)
	}

	for i := 0; i < 13; i++ {
		fmt.Println("~~~~~Loop ", i)
		for k, v := range mymap {
			fmt.Println(k, v)
		}
	}
}
