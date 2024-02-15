package main

import (
	"fmt"
	"time"
)

/*
1. channel에 크기가 설정되어있지 않는 경우
*/

// a. deadlock 현상 발생
//func main() {
//	ch1 := make(chan int)
//	slInt := make([]int, 10)
//
//	go func() {
//		// ch1 에 순차적으로 할당을 해줌
//		for i, _ := range slInt {
//			time.Sleep(1 * time.Second)
//			ch1 <- i + 1
//
//		}
//		//close(ch1) // 대기조로 남아있는 서브루틴으로 인한 deadlock을 방지하기 위해 채널을 닫아준다.
//	}()
//
//	// 대기조로 남아있는다.
//	for val := range ch1 {
//		fmt.Println(val)
//		time.Sleep(time.Second)
//	}
//	fmt.Println("끝이야 ~!")
//}

// b. 성능 이슈 - 블로킹 현상이 일어남에 따른 추가 지연시간
//func main() {
//	c := make(chan struct{})
//	go func() {
//		time.Sleep(2 * time.Second)
//		<-c
//	}()
//	start := time.Now()
//	c <- struct{}{} //  여기서 블로킹 현상이 일어납니다
//	elapsed := time.Since(start)
//	fmt.Printf("Elapsed: %v\n", elapsed) // 블로킹 현상이 일어나는 시간 산출 => ex) Elapsed: 2.0125838s
//}

// c. 리소스 누수 - 무한루프내 채널에 값을 지속적으로 보내는 코드를 위치시킬시 채널의 버퍼는 실시간으로 채워지게 되고, 값을 빼낼 수 없는 상태라 블록이 된다.

func main() {
	ch := make(chan int)
	go func() {
		i := 1
		for {
			ch <- i
			i++
			// 적절한 처리 시간을 주지 않으면 메모리 누수가 발생할 수 있음
			time.Sleep(time.Second * 2)
		}
	}()

	// 1. 종료 케이스
	value := <-ch // 여기서 종료가 될시에는 채널에서 값을 받아오지 않으면 고루틴이 블록되어 더 이상 값을 전송하지 않음 => 종료 됨 -=> 메모리 누수가 발생하지 않는다.
	fmt.Println("Received:", value)

	// 2. 미종료 케이스
	// 아래 주석 처리된 부분을 주석 해제하면 블록을 피하고 메모리 누수가 발생할 수 있음
	// time.Sleep(time.Second * 5)
}
