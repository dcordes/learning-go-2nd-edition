package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(fname string, lname string, age int) Person {
	return Person{
		firstName: fname,
		lastName:  lname,
		age:       age,
	}
}

func MakePersonPointer(fname string, lname string, age int) *Person {
	return &Person{
		firstName: fname,
		lastName:  lname,
		age:       age,
	}
}

func main() {
	john := MakePerson("John", "Jones", 20)
	jill := MakePerson("Jill", "Johnson", 21)
	fmt.Println(john)
	fmt.Println(jill)
}
