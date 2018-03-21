package main

import "fmt"

// define a new type (a struct)
type person struct {
	name string
	age  int
}

// birthday for the person
// try what happens if the functions operates on person
// instead on *person (modify it here)
func (p *person) birthday() {
	p.age++
}

// this struct inherits the fields from person
type employee struct {
	person
	jobtitle string
}

func main() {

	alice := person{name: "Alice", age: 25}

	fmt.Printf("meet %v\n", alice)
	alice.birthday()
	fmt.Printf("meet %v again\n", alice)

	bob := employee{}
	// can access the field directly (inherited from person struct)
	bob.age = 30
	// employee bob can act as a receiver for the birthday call
	bob.birthday()
}
