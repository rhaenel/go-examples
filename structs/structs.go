package main

import "fmt"

// define a new type (a struct)
type person struct {
	name string
	age  int
}

// this function takes a pointer to a struct as a parameter,
// so it can modify the struct
func birthday(p *person) {
	p.age++
}

func main() {

	alice := person{name: "Alice", age: 25}
	bob := person{name: "Bob", age: 31}

	fmt.Printf("%v meets %v again\n", alice, bob)

	birthday(&alice)
	fmt.Printf("%v meets %v again\n", alice, bob)

	// allocates a new struct and returns a pointer
	sean := new(person) // because of the garbage collector, pretty much identical to sean := &person{}
	sean.name = "Sean"
	birthday(sean)
	fmt.Printf("meet %v\n", *sean)

	// dereference of struct fields always uses '.', no matter pointer or not
	fmt.Printf("ages: %d %d %d", alice.age, bob.age, sean.age)
}
