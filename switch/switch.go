package main

import "fmt"

func main() {

	i := 2
	fmt.Print("i = ")
	// basic switch usage
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Printf("%d\n", i)
	}

	j := 2
	// can use multiple values in case statemenet
	switch j {
	case 1: // beware of this
	case 2: // probably not what you want
	case 3:
		fmt.Println("j = 3")
	case 4, 5, 6:
		fmt.Println("j = 4..6")
	}

	a := 20
	b := 1
	switch {
	case a < 10:
		fmt.Println("a < 10")
	case b < 10:
		fmt.Println("b < 10")
	default:
		fmt.Println("default")
	}
}
