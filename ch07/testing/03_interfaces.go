package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Alien struct {
	PlanetAge int
}

// because we are using a POINTER RECEIVER here, the "BumpAge()" method:
//
//	-- IS in the method set of the type called "*Person"
//	-- is NOT in the method set of the type called "Person"
func (p *Person) BumpAge() {
	p.Age++
}

func (p Person) PrintName() {
	fmt.Println(p.FirstName + " " + p.LastName)
}

func (a *Alien) BumpAge() {
	a.PlanetAge = a.PlanetAge + 1000
}

type Creature interface {
	BumpAge()
}

type Being interface {
	PrintName()
}

func main() {
	// testing structs and methods again
	John := Person{
		FirstName: "John",
		LastName:  "James",
		Age:       10,
	}
	Xypz := Alien{
		PlanetAge: 5000,
	}
	fmt.Println(John)
	fmt.Println(Xypz)

	John.BumpAge()
	Xypz.BumpAge()
	fmt.Println(John)
	fmt.Println(Xypz)
	John.PrintName()

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// now turning to interfaces

	fmt.Println("~~~~~~~~~~~")
	fmt.Println("An interface assignment that will only work if the type's POINTER (not value) is passed")

	// (1) declare a variable based on the Creature interface
	// 		-- that is to say, any value can be assigned to this variable so long as it is a type that implements the method set associated with it; the method set available to *Creature* is *BumpAge*

	var myCreature Creature

	// (2) try assigning preexisting variables of various types to that inteface variable
	//		-- if the preexisting variable is of a type that implements the methods in the interface's method set, then this should succeed

	// this first attempted assignement will fail
	//    * error:
	//			cannot use John (variable of type Person) as Creature value in assignment: Person does not implement Creature (method BumpAge has pointer receiver)
	//	  * why?  meaning?
	//    * because the assignment is saying that in the method set of "John" is the "BumpAge()" method
	//	  * but that is NOT true -- the "BumpAge()" method is in the method set of "*John" (not "John") where that method is defined with the "*Person" pointer receiver
	//myCreature = John

	myCreature = &John
	fmt.Println(myCreature)
	myCreature.BumpAge()
	fmt.Println(myCreature)

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	fmt.Println("~~~~~~~~~~~")
	fmt.Println("An interface assignment that will work if EITHER the type's POINTER or VALUE is passed")

	// (1) declare variables based on the Being interface
	// 		-- that is to say, any value can be assigned to these variables so long as it is a type that implements the method set associated with it; the method set available to *Being* is *PrintName*

	var myBeingPtr Being
	var myBeingVal Being

	// (2) try assigning preexisting variables of various types to that inteface variable
	//		-- if the preexisting variable is of a type that implements the methods in the interface's method set, then this should succeed

	// unlike the first example above, here we can assign John (based on the Person type) to a variable of the Being interface
	//  EITHER by passing the pointer of John or the value of John
	// whereas in the previous example it only worked if we passed the pointer reference.  Why?
	//
	// BECAUSE the receiver for the PrintName method is a VALUE receiver (not a pointer receiver), and so the interface will look in the method set
	// 	of "John" for the PrintName method (and not the method set of "*John" as above)
	//
	// So that explains why "myBeingVal = John" does not error out
	//
	// But why doesn't "myBeingPtr = &John" error out?  Because, PRESUMABLY, Go can dereference the pointer and find the value it needs and search there for the method set
	// But the reverse is not true!  If it is expecting a pointer (as in the case above) and you pass it a value, it cannot dereference it
	//	and find the value!
	//
	// So to sum up, you can make a successful assignment of a variable of a given type to a given interface if
	//	1) EITHER the type of that variable supports a VALUE receiver and
	//		a} in the assignment to the interface you specify the POITNER to the variable, OR
	//		b} you provide the actual VALUE of the variable
	//		~~~~
	//	2) OR the type of that variable supports a POINTER receiver and
	//		a) in the assignment to the interface you specify the POINTER to the variable

	myBeingPtr = &John
	fmt.Println(myBeingPtr)
	myBeingPtr.PrintName()

	myBeingVal = John
	fmt.Println(myBeingVal)
	myBeingVal.PrintName()

}
