package main

import "fmt"

func main() {

	// basic usage of if
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	a := 1.0
	b := 2.0
	// can declare local variable in if statement, accessible in all else branches
	if c := a / b; c > 0.5 {
		fmt.Println("bigger half")
	} else if c < 0.5 {
		fmt.Println("smaller half")
	} else {
		fmt.Println("half")
	}
}
