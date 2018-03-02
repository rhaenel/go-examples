package main

import "fmt"

func main() {
	var a [5]int                 // an array of 5 integers
	fmt.Printf("1: a = %v\n", a) // can be formatted with the standard %v

	a[2] = 42                    // change an element
	fmt.Printf("2: a = %v\n", a) // print modified array

	b := [5]int{1, 2, 3, 4, 5} // type of b is inferred, uses specific array initializer
	fmt.Printf("3: b = %v\n", b)

	a = b                        // this effectively copies the array
	a[2] = 0                     // modify some element
	fmt.Printf("4: a = %v\n", a) // this should show the modified version
	fmt.Printf("5: b = %v\n", b) // this should show the unmodified version

	var c [10]int
	//c = a // try this: it won't work, because [5]int and [10]int are different types!

	_ = c // special assignment to prevent the "declared and not used" error for c
}
