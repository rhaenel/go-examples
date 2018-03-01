package main

import "math"

func main() {
	// 'const' can be used everywhere where 'var' appear,
	// however there is no short form like :=

	const c = 3 // special case: type of constant not defined yet

	const k int = 3 // constant with type inferred as int16

	const f float64 = 3 // constant with fixed type float64

	math.Sin(c) // this works because the c is auto-cast to float64 (Sin needs float64)

	math.Sin(k)          // this doesn't work because k is of type int
	math.Sin(float64(k)) // and must be manually cast to float64 to be used here
}
