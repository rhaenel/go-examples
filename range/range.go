package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}

	// iterate the elements in a slice. looks the same for arrays
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range provides the index for each entry. In the previous
	// example this was ignored by using the _ variable. Here
	// we use it
	for i, num := range nums {
		fmt.Printf("index %d num %d\n", i, num)
	}

	fruits := map[string]int{"strawberries": 11, "raspberries": 10}

	// range on maps delivers key, value pairs
	for k, v := range fruits {
		fmt.Printf("%v -> %v\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range fruits {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// Note that the length of the string counts the bytes
	// of the constant's UTF-8 representation
	s := "Köln"
	fmt.Printf("string length: %d\n", len(s))
	for i, c := range s {
		fmt.Printf("string index %d: %#U\n", i, c)
	}
}
