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
			fmt.Println("go-1-할당전 : ", i)
			ch <- i // 1. 자리가 있는지 여부를 확인하고 꽉차있으면 대기합니다/  3. 자리가 있으면 값을 할당하게 됩니다.
			fmt.Println("go-1-할당후 : ", i)
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
			test := <-ch // 2. 값을 꺼낸뒤에 자리를 만들게 되면
			fmt.Println("go-2-값꺼냄 : ", test)
			time.Sleep(5 * time.Second)
		}
	}()

	wg.Wait()
}
