package main

import "fmt"

type oscar interface {
	CheapShot()
}

func main() {
	var teddy oscar

	if teddy == nil {
		fmt.Println("It's nil! -- Which means that this should (and does!) panic...")
		fmt.Println("And why is it nil?  Because although it has a TYPE set it does not have a VALUE set")

		teddy.CheapShot()
	}
}
