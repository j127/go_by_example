package main

import (
	"fmt"
	"time"
)

func main() {
	// string chan with 1 slot
	channel1 := make(chan string, 1)

	go func() {
		time.Sleep(2200 * time.Millisecond)
		channel1 <- "here's result 1"
	}()

	select {
	case res := <-channel1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	fmt.Println("done")
}
