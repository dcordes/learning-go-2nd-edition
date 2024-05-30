package main

import "fmt"

const (
	MailCategory = 10
	Personal
	Spam
	Social = 30
	Advertisements
)

func main() {
	fmt.Println(MailCategory, Personal, Spam, Social, Advertisements)
}
