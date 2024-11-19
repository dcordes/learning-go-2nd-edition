# Chapter 03

## Maps

* best to work with EMPTY maps instead of NIL maps

* e.g. EMPTY MAP LITERAL
	* `scores := map[int][]string{}`
* e.g. NIL MAP
	* `var scores map[int][]string`

* detailed example:
```
package main

import "fmt"

func main() {

	// this works! --> here we are using what is called an EMPTY MAP LITERAL
	scores := map[int][]string{}

	// but this wouldn't have worked, in fact would cause a panic:
	//	"panic: assignment to entry in nil map"
	// because here we are using a NIL MAP
	// writing to a nil map, as we do in the assignment line with "fred" CAUSES A PANIC
	//
	//var scores map[int][]string

	scores[0] = append(scores[0], "fred")

	fmt.Println(scores)
}
```

* example of how to get all keys in a map into a slice

```
	for name := range l.Wins {
		names = append(names, name)
	}
```