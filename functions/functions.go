package main

import "fmt"

// add to ints
func plus(a int, b int) int {

	return a + b
}

// parameters of the same type can be grouped
func plusPlus(a, b, c int) int {
	return a + b + c
}

// functions can return multiple values
func plusMinus(a, b int) (int, int) {
	return a + b, a - b
}

// variable length argument list
func sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func main() {

	// calling is just like you expect
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	p, m := plusMinus(1, 2)
	fmt.Printf("1 +- 2 = %d, %d\n", p, m)

	fmt.Printf("1+2+3+4+5 = %d\n", sum(1, 2, 3, 4, 5))

	// anonymous functions used as a type
	f := func(i int) int {
		return i * 2
	}
	fmt.Printf("calling anonymous function: %d\n", f(2))
}
