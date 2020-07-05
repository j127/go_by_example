package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // 0 because it wasn't initialized

	// shorthand for declare and initialize
	// same as `var f string = "apple"`
	f := "apple"
	fmt.Println(f)
}
