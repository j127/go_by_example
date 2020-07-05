package main

import "fmt"

type hobbit struct {
	name string
	age  int
}

func newHobbit(name string) *hobbit {
	h := hobbit{name: name}
	h.age = 42

	// You can safely return a pointer to local variable as a local variable
	// will survive the scope of the function.
	return &h
}

func main() {
	fmt.Println(hobbit{"Alice", 100})
	fmt.Println(hobbit{name: "Bob", age: 100})

	// ommitted fields will be zero-valued
	fmt.Println(hobbit{name: "Merry"}) // age 0

	fmt.Println(newHobbit("Pippin"))

	fmt.Println(hobbit{"Bilbo", 111})
	fmt.Println(&hobbit{"Frodo", 33})

	s := hobbit{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	// automatically dereferenced
	fmt.Println(sp.age) // 50

	sp.age = 51
	fmt.Println(sp.age) // 51
}
