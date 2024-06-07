# Chapter 07

## Types in Go

* you can declare not just struct types but other types

```
type Person struct {
  FirstName string
  LastName  string
  Age       int
}
type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```
* ~~~~~~~~

* types can only be accessed from within their scope
* see (!!!) below for a comment on the `type Score int` example above

* ~~~~~~~~

* and here is an example of using and defining a string slice in a struct
* note that you have to re-specify the string slice type in the definition (otherwise you will get a `missing type in composite literal` error)

```
type Team struct {
	team_name    string
	player_names []string
}

aaa := Team{
	team_name:    "The AAA Team",
	player_names: []string{"A", "AA", "AAA"},
}
```

## Methods

* the "functions" associated with a type are called "methods", just like the "functions" associated with "objects" in other languages
* method declarations look like regular function declarations, with the addition of the *receiver*; the receiver is what types the method to type
  * the receiver is specified between the keyword *func* and the method name
  * in this case the method name is "String"; it takes no arguments and returns a striing
  * the receiver itself is `(p Person)`
    * the "Person" reference ties it to the Person type
    * the name of the receiver itself in this case is "p"
      * it is typical to use the first lower-case letter of the type; but (apparently) `this` or `self` can also be used

```
type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func (p Person) String() string {
  return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```

* methods for a type are defined at the package block level *only*, unlike functions which can be declared in any block
* method names cannot be overloaded within the same type
* methods must be declared in the same package as their associated type
* method invocations are standard:

```
p := Person{
  FirstName: "Fred",
  LastName:  "Fredson",
  Age:       52,
}

output := p.String()
```

### Pointer Receivers and Value Receivers

* methods (i.e. the functions associated with a type) can have both pointer receivers (see the "Increment" example here) and value receivers (see the "String" example here)
  * you *have* to use a pointer receiver if your method modifies the receiver (of course) or if the method needs to handle *nil* instances (see "Code your methods for nil instances below")
  * otherwise you *can* use a value receiver... but it is common practice to use *all* pointer receivers for a type's method if even *one* of the type's method has a pointer receiver

```
type Counter struct {
    total       int
    lastUpdated time.Time
}
func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
```

You could call the above thus:

```
var c Counter
fmt.Println(c.String())
c.Increment()
fmt.Println(c.String())
```

Notice that although "Increment" is a pointer receiver method that it is being called against a "value" variable, "c", not a pointer.  In these cases Go automatically takes the address, so `c.Increment()` is auto-converted to `(&c).Increment()`.

Ditto the reverse.  In the below `c.String()` gets dereferenced: `(*c).String()`.

```
c := &Counter{}
fmt.Println(c.String())
c.Increment()
fmt.Println(c.String())
```

* Go *discourages* writing getter/setter methods and instead encourages directly accessing the fields of a type.

### Code Your Methods for nil Instances

* if your method has a pointer receiver it should explicitly handle what to do if it receives a *nil* pointer; mainly this means one of these options:
  * panic
  * return an error

### Methods are Functions Too

* you can work with a type via a function, as in "FuncBumpUp" below, or you can use a method to tie that work directly to the type, as in "BumpUp" below; both of these works, but the more direct method tie is preferable

```
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
```

* you can also have a value that is a method, a *method value*; see in this example how variable "c" is declared...
  * first with a variable of a given type
  * then with a method of that type
* ... such that it can then be invoked as `c("Lord")` instead of `catan.FancyName("Majesty")` as earlier

```
func (g Game) FancyName(prefix string) string {
	return prefix + " " + g.Name
}

func main() {
	catan := Game{
		Name:   "Settlers",
		Rating: 5,
	}
  fmt.Println(catan.FancyName("Majesty"))
	c := catan.FancyName
	fmt.Println(c("Lord"))
}
```

* you can also create a `method expression` (skipping this explanation)

### Functions Versus Methods

* use methods when your operations are dependent on data defined at startup or an earlier operations
* use functions, on the other hand, if an operation can be performed purely on the basis of input parameters

### Type Declarations Aren't Inheritance

* although a custom type can be defined based on a custom type, *there is no inheritance of methods in Go*
  * (maybe interfaces allow you to compensate for lack of method sharing that inheritance provides?  no -- it's via *embeddeding* -- see below

### Types Are Executable Documentation

* using structs makes code much more readable
* in the same way, it makes sense to base one type off of another custom type if both types share the same data fields, but if you want to have different methods on them

## iota is for Enumerations -- Sometimes

* reminder on how constant value assignments work -- if a value is not specified the value gives come from the previous assignment; so the below prints out `10 10 10 30 30 30`

```
const (
	MailCategory = 10
	Personal
	Spam
	Social = 30
	Advertisements
	Ox
)

func main() {
	fmt.Println(MailCategory, Personal, Spam, Social, Advertisements, Ox)
}
```

* add the `iota` keywork into the mix:
  * within the const block a background counter starts at 0 and increments for each field in the const block
  * whenever `iota` is invoked the background counter value is displayed
    * or if the value is unspecified, and iota was the last articulated value

```
const (
	MailCategory1 = 10
	Personal1     = iota
	Spam1
	Social1 = 30
	Advertisements1
	Ox1 = iota
)

func main() {
	fmt.Println(MailCategory1, Personal1, Spam1, Social1, Advertisements1, Ox1)
}
```

* when to use `iota`?  Bohner: "iota-based enumerations make sense only when you care about being able to differentiate between a set of values and don't particularly care what the value is behind the scenes"
  * i.e. if you *do* care about particular values then don't use it because the value assigned to a variable will change in you insert a new value into the middle of the list

## Use Embedding for Composition

* now although out-and-out inheritance is not supported, *embedding* is; note how in the example below the *Description* method can be invoked on a varaible instantiated from the Manager struct, even though that Description method is not part of that struct; why?  because Manager has Employee embedded within it with *no name assigned to the field* -- that makes it an *embedded field*

```
type Employee struct {
    Name         string
    ID           string
}
func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}
type Manager struct {
    Employee
    Reports []Employee
}
func (m Manager) FindNewEmployees() []Employee {
      // do business logic
}

m := Manager{
    Employee: Employee{
        Name: "Bob Bobson",
        ID:   "12345",
    },
    Reports: []Employee{},
}
fmt.Println(m.ID)
fmt.Println(m.Description()) // prints Bob Bobson (12345)
```

* note in the above that when the Description  method was invoked that no reference was made to Employee; that works fine; the only time you have to be more explicit is when you have field or method names overlap, as in this example where it is necessary to say `o.Inner.X`:

```
type Inner struct {
    X int
}
type Outer struct {
    Inner
    X int
}

o := Outer{
    Inner: Inner{
      X: 10,
    },
    X: 20,
}
fmt.Println(o.X)
fmt.Println(o.Inner.X) // prints 10
```

## Embedding Is Not Inheritance

## A Quick Lesson on Interfaces

* his example here is terrible
* to note, "String" is a value receiver and "Increment" is a pointer receiver
* see my "03_interfaces.go" file for an explanation
  * my DETAILED long explanation
* ~~~~
* now, why use interfaces at all?
  * if you want to perform a common operation (e.g. GetName) for a variety of different types, all of which support that operation, using just one function
  * that function must specify the type of the input parameter, and since you are dealing with different types, you need a way of referring to all the different types you can support
  * that's what the interface does -- defines all the different types that have that function
  * ~~~~~
  * what about a function that we want to work with ANY type?
    * like fmt.Println()
    * I think that's what the Empty interface is for -- see below...

## Interfaces Are Type-Safe Duck Typing

* *Design Patterns* says "Program to an interface, not an implementation"
* Go's interfaces are implemented *implicitly*
  * that is, we don't declare anywhere that a particular struct implements an given interface
  * it's just that if that struct has all of the methods in the interface's method set, that it implements it
* it is of course fine for a type to have additional methods not in those of the interface it implements

## Embedding and Interfaces

* you can embed not only structs in structs, but also interfaces in interfaces, e.g.:

```
type Reader interface {
        Read(p []byte) (n int, err error)
}
type Closer interface {
        Close() error
}
type ReadCloser interface {
        Reader
        Closer
}
```

## Accept Interfaces, Return Structs

* Go developers say, "Accept interfaces, return structs"
* Why?  Because if you return a concrete type (i.e. struct) you can add methods and fields to it over time without breaking clients (since they will ignore the new fields/methods); but if you return an interface and then subsequently add a method to it then the clients will break (as the struct they are working with that implements that interface may not support that new method)
* try to write separate factory functions for each concrete type
* of course when errors are returned they are often of the interface type
* ~~~
* interfaces are the only abstract type in go

## Interfaces and nil

* in Go interfaces are implemented as structs with two pointer fields: (a) one for the value and (b) one for the type of the value
* an interface is only equal to *nil* if BOTH of these fields are equal to nil
* if an interface variable is nil then invoking any methods on it generates a panic
* see the 04_... file in the _testing_ directory that shows this`

## Interfaces are Comparable

* recall you can declare a type that isn't a struct!
  * (link, !!!, to the original comment on this above)
* and that non-struct type can also have methods on it!
* example of this (see example 05 in testing folder)

```
package main

import "fmt"

type basicint int

func (b basicint) Triple() {
	fmt.Println(b * 3)
}

func main() {
	bi := basicint(6)

	fmt.Println(bi)

	bi.Triple()
}
```

* ~~~~~~~~~~~~~~
* two interfaces are equal if their TYPE and VALUE fields are both equal
* keep in mind that slices are not comparable

## The Empty Interface Says Nothing

* the empty interface is an interface with no associated method signatures; e.g.

```
type oscar interface{}
```

* as opposed to:

```
type oscar interface {
	BumpAge()
}
```

* an empty interface can thus store a variable of any type:

```
	var empty interface{}

	empty = 5
	empty = "test"
	fmt.Println(empty)
```

* `any` can now be used as a substitute for `interface{}`

```
	var empty any
```

* empty types are often used when reading data of unknown types, like reading from JSON
* try to avoid using them -- unidiomatic -- Go is intended to be a strongly typed language!

## Type Assertions and Type Switches

* there are two ways to confirm the type of a variable *or to confirm if a variable implements some other interface*
  * (1) *Type Assertion*: here you *assert* the type to a variable and then check for errors
  * (2) *Type Switches*: here you ask for the type of the variable, and then perform different actions, often in a Switch statement, based on what is returned
* example of former

```
	type abc any
	var first_var abc = 10
	var secon_var abc = "test"
	// imagine that the assignment of first and second var above took place elsewhere, and we want to perform an operation on them, but only if the types are right
	my1, ok := first_var.(int)
	if !ok {
		fmt.Println("incompatible type for first var")
	} else {
		fmt.Println(my1 + 2)
	}
	my2, ok := secon_var.(int)
	if !ok {
		fmt.Println("incompatible type for first var")
	} else {
		fmt.Println(my2 + 2)
	}
  ```

* example of latter

```
	my_type_switch := func(i any) {
		switch j := i.(type) {
		case nil:
			//
		case int:
			fmt.Println(j)
			fmt.Println(i.(int) + 9)
		case string:
			fmt.Println(j)
			fmt.Println(i.(string) + " hello")
		default:
			fmt.Println(j)
		}
	}

	my_type_switch(first_var)
	my_type_switch(secon_var)
```

* a better way to do this is to use "reflection" -- see chapter 16

* now here is another example that shows how you can try to see if a variable implements one interface or another, and take different action based on what you find

```
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
```

## Use Type Assertions and Type Switches Sparingly

* when do we use type assertions and switches?  not too often -- sometimes when we want to determine the concrete type of an interface so that we can then go and see if it implements other interfaces
  * I gave an example of some custom code that does this in the previous section
* for example, involving an API with the use of "context" (covered in chapter 14):
  * if the input variable supports some interfaces, do this, otherwise do that

* when determining what type of error you have received (recall that errors implement the `error` interface), use `errors.Is` and `errors.As` to check error types

## Function Types Are a Bridge to Interfaces

* ofen you have a choice to specify a function as an input parameter to a function, or to attach that function as a method...
  * (kind of skimmed this section)

## Implicit Interfaces Make Dependency Injection Easier

* (also skimmed this section)
* remember "accept interfaces, return structs"?  another way to same the reason for that is that it helps with "dependency injection" -- meaning that
  * our functions accept interfaces rather than explicit parameters
  * and those interfaces of course describe the parameters that those objects must support
  * the definition of the interface of course is defined elsewhere
  * so in a sense the dependencies of the function, what it needs the parameters it receives to be able to do, are not specified in the function itself, but elsewhere, meaning that they get *injected* into the function
  * this is the type of dependency injection called "interface injection" (there are two other types)
  * some links:
    * https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6
    * https://www.freecodecamp.org/news/a-quick-intro-to-dependency-injection-what-it-is-and-when-to-use-it-7578c84fa88f/
    * https://martinfowler.com/articles/injection.html
* we use interfaces all the time in unit testing

## Wire

* Wire is a tool from Google that helps with dependency injection:
  * https://github.com/google/wire
* I didn't see any references to this in our Go codebase

## Go Isn't Particularly Object-Oriented (and That's Great)

* personal note from D. Cordes here...
* one consequence of Go not being object-oriented is that there is not a native concept of constructors
* but sometimes you need to rig something up
* e.g., look at the `if l.Wins == nil` line below; if this is omitted then if you try to populate the map with a command like `l.Wins[first_team_name] = first_team_score` then you will receive a `panic: assignment to entry in nil map` error

```
type Team struct {
	team_name    string
	player_names []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(
	// using a pointer receiver here because we are modifying the elements of the receiver
	first_team_name string,
	first_team_score int,
	second_team_name string,
	second_team_score int,
) {
	// if the map has not yet been initialized, initialize it
	if l.Wins == nil {
		l.Wins = make(map[string]int)
	}
	// populate with the match results
	l.Wins[first_team_name] = first_team_score
	l.Wins[second_team_name] = second_team_score
}
```

## Some questions to ask the book club about this chapter

* pointer/value receiver confusion in exer 02
* lack of constructors in structs since Golang not OOO
	* corresponding need to initialize things like maps that are part of structs
* dependency injection (and at Datadog)
	* no references to wire (i.e. "InitializeEvent") in dd-go
