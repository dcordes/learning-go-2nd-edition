package main

import "fmt"

type Game struct {
	Name   string
	Rating int
}

// a function that works with that type (but is not a method)
func FuncBumpUp(g *Game) {
	g.Rating++
}

// a pointer receiver method, now tied directly to the type with the receiver
func (g *Game) BumpUp() {
	g.Rating++
}

// a value receiver method
func (g Game) FancyName(prefix string) string {
	return prefix + " " + g.Name
}

func main() {
	catan := Game{
		Name:   "Settlers",
		Rating: 5,
	}
	fmt.Println(catan)
	FuncBumpUp(&catan)
	catan.BumpUp()
	fmt.Println(catan)
	fmt.Println(catan.FancyName("Majesty"))

	fmt.Println("~~~ method value")
	c := catan.FancyName
	fmt.Println(c("Lord"))

}
