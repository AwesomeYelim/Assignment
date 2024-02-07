package main

import (
	"fmt"
	"time"
)

// 채널의 데이터 전송 방법은 기본적으로 FIFO(선입선출) 대기열인 채널의 동작을 기반함
func main() {
	//var wg sync.WaitGroup

	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	// ch1 에 순차적으로 할당을 해줌
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	close(ch1)

	for val := range ch1 {
		ch2 <- val
		//defer wg.Done()
		//time.Sleep(time.Second)
		fmt.Println("다른작업")
		time.Sleep(time.Second)
	}
	close(ch2)

	// 루프를 돌려보면 할당한 순서대로 입력이 되는걸 볼수 있다.
	for val := range ch2 {
		//time.Sleep(time.Second)

		fmt.Println(val)
	}

	//wg.Wait()
}
