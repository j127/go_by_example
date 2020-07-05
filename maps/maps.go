package main

import "fmt"

func main() {
	// create an empty map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("value from k1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// does `prs` stand for present?
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	val, prs2 := m["k1"]
	fmt.Println("prs2:", prs2)
	fmt.Println("value from prs2:", val)

	// assign values while declaring the map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
