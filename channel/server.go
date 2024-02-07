package main

import (
	"fmt"
	"net"
	"sync"
)

type Message struct {
	Text string
}

func main() {
	var wg sync.WaitGroup

	listener, _ := net.Listen("tcp", "localhost:3000")           // tcp 서버 생성
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001") // udp 생성

	defer listener.Close()

	// a,b channel 을 생성
	AChannel := make(chan Message, 2) // 초기 버퍼를 설정하는 이유는 => 처음 채널 크기는 0이므로 데이터를 빼갈때 까지 대기함 => 데이터를 가져가지 않아서 프로그램이 멈추는 현상이 생긴다(deadlock)
	BChannel := make(chan Message, 2)

	udpconn, err := net.DialUDP("udp", nil, serverAddr) // udp 패킷에서는 정보를 어디서 먼저 줄것인지를 표시하기위해 Dial을 먼저 작성해준다.
	if err != nil {
		fmt.Println(nil)
		return
	}
	tcpconn, err := listener.Accept()
	if err != nil {
		fmt.Println(nil)
		return
	}

	defer tcpconn.Close()
	defer udpconn.Close()

	wg.Add(1)
	// 통신을 받아서 a channel 에 할당하는 함수
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _ := tcpconn.Read(buffer)
			AChannel <- Message{Text: string(buffer[:n])}
		}
	}()

	wg.Add(1)
	// b channel에 할당
	go func() {
		for {
			BChannel <- <-AChannel
		}
	}()

	wg.Add(1)
	// udp 연결루틴 생성 및 송신
	go func() {
		defer wg.Done()
		for {
			response := <-BChannel
			// 클라이언트로 msg write
			msg := []byte(response.Text)
			_, _ = udpconn.Write(msg)
			fmt.Println(string(msg))
		}
	}()
	wg.Wait()
	//close(BChannel)
}
