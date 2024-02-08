package main

import (
	"fmt"
	"time"
)

/*
1. channel에 용량이 설정되어있지 않는 경우
*/

// a. channel값 사용 => go routine 사용한 경우
func main() {
	//var wg sync.WaitGroup

	ch1 := make(chan int)
	slInt := make([]int, 10)

	//wg.Add(1)
	// 대기조로 남아있는다.
	go func() {
		//defer wg.Done()
		for val := range ch1 {
			fmt.Println(val)
			time.Sleep(time.Second)
			//if val == 7 {
			//	time.Sleep(time.Second * 10)
			//}
		}
		fmt.Println("끝이야 ~!")
	}()

	// ch1 에 순차적으로 할당을 해줌
	for i, _ := range slInt {
		ch1 <- i + 1
		//fmt.Println(<-ch1)
	}
	close(ch1) // 대기조로 남아있는 서브루틴으로 인한 deadlock을 방지하기 위해 채널을 닫아준다.
	time.Sleep(5 * time.Second)
}

// b. channel에 값 할당 => go routine 사용한 경우
//func main() {
//	// 비버퍼 채널 생성
//	ch := make(chan int)
//	slInt := make([]int, 10)
//
//	go func() {
//		for i, _ := range slInt {
//			ch <- i + 1
//		}
//		close(ch)
//	}()
//
//	for value := range ch {
//		fmt.Println(value)
//		time.Sleep(time.Second)
//	}
//	fmt.Println("끝이야 ~!")
//}
