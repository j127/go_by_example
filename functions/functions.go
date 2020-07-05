package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// you can omit the types until the last item if they are the same
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
