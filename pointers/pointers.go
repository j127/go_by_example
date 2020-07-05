package main

import "fmt"

// args get passed by value
func zeroval(ival int) {
	ival = 0
}

// args get passed by reference
func zeroptr(iptr *int) {
	// set value of dereferenced pointer
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// pass in the memory address of (pointer to) `i`
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("it's still zero:", i)

	fmt.Println("pointer:", &i)
}
