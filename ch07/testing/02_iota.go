package main

import "fmt"

const (
	MailCategory = 10
	Personal
	Spam
	Social = 30
	Advertisements
	Ox
)
const (
	MailCategory1 = 10
	Personal1     = iota
	Spam1
	Social1 = 30
	Advertisements1
	Ox1 = iota
)

func main() {
	fmt.Println(MailCategory, Personal, Spam, Social, Advertisements, Ox)
	fmt.Println(MailCategory1, Personal1, Spam1, Social1, Advertisements1, Ox1)
}
