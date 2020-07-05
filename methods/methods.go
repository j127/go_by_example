package main

import "fmt"

type rect struct {
	width, height int
}

// a _receiver_ type of `*rect`
func (r *rect) area() int {
	return r.width * r.height
}

// this one has a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// define a rectangle
	r := rect{width: 10, height: 5}

	// call the methods
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for
	// method calls. You may want to use a pointer receiver type to avoid
	// copying on method calls or to allow the method to mutate the receiving
	// struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
