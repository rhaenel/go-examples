package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // the basic form (well-known from C)
		fmt.Printf("i = %d\n", i)
	}

	j := 5
	for j > 0 { // 'for' is also 'while' in Go!
		fmt.Printf("j = %d\n", j)
		j--
	}

	n := 3
	for { // 'for' is also 'while(true)' in Go!
		fmt.Printf("n = %d\n", n)
		if n >= 3 {
			n--
			continue // on with the loop
		}
		break // the only way to get out of here
	}

outerloop: // loops can be named to address them in a break/continue
	for {
		for m := 0; m < 10; m++ {
			fmt.Printf("m = %d\n", m)
			if m == 3 {
				break outerloop
			}
		}
	}
}
