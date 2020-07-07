package main

// This is a quick script to generate one file of notes for each chapter.
// I will fill in the chapter topic from memory and then use active
// recall to try to reinforce the information.

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func writeFile(filename string) {
	text := []byte(filename)
	fn := fmt.Sprintf("%s.md", filename)
	ioutil.WriteFile(fn, text, 0644)
}

// My peg list runs from 200-274.
func main() {
	fmt.Println("beginning")
	for i := 200; i <= 274; i++ {
		writeFile(strconv.Itoa(i))
	}
	fmt.Println("done")
}
