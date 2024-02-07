package main

import "fmt"

// 채널의 데이터 전송 방법은 기본적으로 FIFO(선입선출) 대기열인 채널의 동작을 기반함
func main() {

	ch1 := make(chan int, 3)

	// ch1 에 순차적으로 할당을 해줌
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) // 값을 더이상 채널로 전송되지 않을때 사용

	for val := range ch1 {
		fmt.Println(val)
	}

}
