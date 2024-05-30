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

* types can only be accessed from within their scope

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
  * (maybe interfaces allow you to compensate for lack of method sharing that inheritance provides?)

### Types Are Executable Documentation

* using structs makes code much more readable
* in the same way, it makes sense to base one type off of another custom type if both types share the same data fields, but if you want to have different methods on them

## iota is for Enumerations -- Sometimes

* reminder on how constant value assignments work -- if a value is not specified the value gives come from the previous assignment; so the below prints out `10 10 10 30 30`

```
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
```

* 

```
const (
    Uncategorized MailCategory = iota
    Personal
    Spam
    Social
    Advertisements
)
```

## Use Embedding for Composition

## Embedding Is Not Inheritance

## A Quick Lesson on Interfaces

## Interfaces Are Type-Safe Duck Typing

## Embedding and Interfaces

## Accept Interfaces, Return Structs

## Interfaces and nil

## Interfaces are Comparable

## The Empty Interface Says Nothing

## Type Assertions and Type Switches

## Use Type Assertions and Type Switches Sparingly

## Function Types Are a Bridge to Interfaces

## Implicit Interfaces Make Dependency Injection Easier

## Wire

## Go Isn't Particularly Object-Oriented (and That's Great)


