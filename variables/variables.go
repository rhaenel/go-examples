package main

import "fmt"

func main() {
	var i int                 // i is of type int, initialized with 0 (default)
	fmt.Printf("i = %d\n", i) // use %d format for int

	var j, k uint32                      // j and k are of type int, initialized with 0 (default)
	fmt.Printf("j = %d, k = %d\n", j, k) //

	var l = 32                // type of l is inferred from the constant (int)
	fmt.Printf("l = %d\n", l) //

	m := 4.2                  // short form with type inference (float in this case)
	fmt.Printf("m = %f\n", m) // use %f because we know it's a float

	n, o := 3, 4.2                       // multiple variable definition with type inference (different types)
	fmt.Printf("n = %v, o = %v\n", n, o) // notice we use %v here, which uses the 'default' format

	s := "I'm a string"  // basic type string is immutable
	fmt.Printf(s + "\n") // + operator effectively creates a new string
}
