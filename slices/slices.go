package main

import "fmt"

func main() {
	a := make([]int, 5)          // make a new slice, backed by an array of 5 int's
	fmt.Printf("1: a = %v\n", a) // feels like an array in Printf

	b := []int{1, 2, 3, 4, 5}    // this also looks much like an array (but it's a slice)
	fmt.Printf("2: b = %v\n", b) // print the slice

	fmt.Printf("3: start  = %v\n", b[:3])  // only the first 3 elements
	fmt.Printf("4: middle = %v\n", b[1:4]) // only the middle 3 elements
	fmt.Printf("5: end    = %v\n", b[2:])  // only the last 3 elements

	c := b                       // this copies the slice (which in turn contains a pointer to the underlying array)
	b[2] = 0                     // modify an element
	fmt.Printf("6: b = %v\n", b) // this contains the modified element
	fmt.Printf("7: c = %v\n", c) // this contains the same modified element!

	d := append(b, 99)           // returns a new slice. this may or may not have a new backing array
	fmt.Printf("8: d = %v\n", d) // print enlarged slice

	copy(a[1:], d[3:])           // copy some slice elements from d to a
	fmt.Printf("9: a = %v\n", a) // show resulting a

	fmt.Printf("10: len(a) = %d\n", len(a)) // len() returns the length of a slice (same for arrays)
}
