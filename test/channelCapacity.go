package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	ch := make(chan int, 2)

	go square()
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Never Print", <-ch)
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
