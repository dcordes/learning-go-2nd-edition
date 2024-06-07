package main

import "fmt"

type iface1 interface {
	CrazyOne()
	CrazyTwo()
}

type iface2 interface {
	CrazyOne()
}

type Tool struct {
	kind string
}

func (t *Tool) CrazyOne() {
	fmt.Println(t.kind + " one")
}

func (t *Tool) CrazyTwo() {
	fmt.Println(t.kind + " two")
}

type Vehicle struct {
	kind string
}

func (v *Vehicle) CrazyOne() {
	fmt.Println(v.kind + " one")
}

func ImplementCrazies(c any) {
	if crazy, ok := c.(iface1); ok {
		fmt.Println("---> Implements interface 1!  Calling both crazy methods")
		crazy.CrazyOne()
		crazy.CrazyTwo()
	} else if crazy, ok := c.(iface2); ok {
		fmt.Println("---> Implements interface 2!  Calling just one crazy method")
		crazy.CrazyOne()
	}
}

func main() {
	my_tool := Tool{
		kind: "wrench",
	}
	my_car := Vehicle{
		kind: "civic",
	}
	ImplementCrazies(&my_tool)
	ImplementCrazies(&my_car)
}
