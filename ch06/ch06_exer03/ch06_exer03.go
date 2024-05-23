package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

// The purpose of this program is to see the effect of changing the GOGC environment
// variable, which controls when garbage collections occurs.  (See README.)  For
// example, at the default value of 100 the below program sees 24 garbage collections, but
// with a value of 800 I see only 4.  (This can be see by setting `GODEBUG=gctrace=1`.)

func main() {
	persons := []Person{}
	for i := 0; i < 10_000_000; i++ {
		persons = append(persons, Person{
			firstName: "Abraham",
			lastName:  "Lincoln",
			age:       50,
		})
	}
}

// Now, if instead of creating just an empty slice of "persons" above you
// instead declare ahead of time what its capacity is...

// 	persons := make([]Person, 0, 10_000_000)

// ... then you will notice that the runtime DRAMATICALLY improves.
