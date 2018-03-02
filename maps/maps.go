package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["apples"] = 0
	m["oranges"] = 3
	m["peaches"] = 1

	fmt.Printf("1: %v\n", m)

	apples, present := m["apples"] // second return values signales if the key is present
	fmt.Printf("2: number of apples: %d, apples are present? %v\n", apples, present)

	delete(m, "apples")      // no more apples
	fmt.Printf("2: %v\n", m) // they should be gone

	apples, present = m["apples"]
	fmt.Printf("4: number of apples: %d, apples are present? %v\n", apples, present)

	l := map[string]int{"strawberries": 11, "raspberries": 10} // initializer short form
	_ = l                                                      // prevent unused... error
}
