package main

import "fmt"
import "math"

// this interface applies to all structs that have both
// methods defined with the struct as a receiver
type geometry interface {
	area() float64
	perim() float64
}

// for example, a rectangle and a circle
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// to implement the `geometry` interface, it is sufficient
// to implement all methods within it, that is, the area()
// and perim() methods in our case
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implementations for the circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a struct satiesfies an interface definition, it
// can be used as an interface type. So a function can
// operate on the interface type (not knowing what
// actual struct is behind it)
func measure(g geometry) {
	fmt.Printf("%v, area %f, perimeter %f\n", g, g.area(), g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// since both circle and rect satisfy the geometry interface
	// we can use them to a call of measure(...) here
	measure(r)
	measure(c)
}
