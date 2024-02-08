package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	size := 3
	ch := make(chan int, size)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		var i int
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("1-1 : ", i)
			ch <- i
			fmt.Println("1-2 : ", i)
			i++
			if size == len(ch) {
				fmt.Println("행행 : ", len(ch))
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			test := <-ch
			fmt.Println("2-1 : ", test)
			time.Sleep(2 * time.Second)
			fmt.Println("2-2 : ", test)
		}
	}()

	wg.Wait()
}
