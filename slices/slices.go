package main

import "fmt"

// slices are typed by their elements, not by length
func main() {
	// an empty slice with non-zero length
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// variable length
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied:", c)

	l := s[2:5]
	fmt.Println("slice1:", l)

	l = s[:5]
	fmt.Println("slice2:", l)

	l = s[2:]
	fmt.Println("slice3:", l)

	// initialize a slice and set the values
	// notice that the length isn't specified in the square brackets
	t := []string{"g", "h", "i"}
	fmt.Println("declared:", t)

	// the inner slices can be different lengths (unlike arrays)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
