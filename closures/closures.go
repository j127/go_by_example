package main

import "fmt"

// function that returns a function
func intSeq() func() int {
	i := 0

	// anon function
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique to each invoked function
	newInts := intSeq()
	fmt.Println(newInts())
}
