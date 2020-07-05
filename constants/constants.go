package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// `math.Sin` expects a `float64` so it gets cast there?
	fmt.Println(math.Sin(n))
}
