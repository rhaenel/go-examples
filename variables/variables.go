package main

import "fmt"

func main() {
	var i int
	i = 1
	fmt.Printf("i has value %d\n", i)

	// type of j will be inferred from int constant
	j := 2
	fmt.Printf("j has value %d\n", j)
}
