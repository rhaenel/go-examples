package main

import "fmt"

func main() {
	var i int

	i = 42
	fmt.Printf("i = %d\n", i)

	var pi *int
	pi = &i
	// show the pointer address and value of the int it points to
	fmt.Printf("pi = %v, pi deref = %v\n", pi, *pi)

	// modify the int through the pointer and show it
	*pi = *pi + 1
	fmt.Printf("i = %d\n", i)
}
