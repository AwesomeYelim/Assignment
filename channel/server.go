package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Message struct {
	Text string
}

func main() {

	var wg sync.WaitGroup

	listenAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1234") // tcp 서버 생성
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001") // udp 생성

	// a,b channel 을 생성
	AChannel := make(chan Message)
	BChannel := make(chan Message)

	listener, err := net.ListenTCP("tcp", listenAddr)

	if err != nil {
		fmt.Println(nil)
		return
	}

	udpconn, err := net.DialUDP("udp", nil, serverAddr) // udp 패킷에서는 정보를 어디서 먼저 줄것인지를 표시하기위해 Dial을 먼저 작성해준다.
	if err != nil {
		fmt.Println(nil)
		return
	}
	defer listener.Close()
	defer udpconn.Close()

	wg.Add(1)
	// 통신을 받아서 a channel 에 할당하는 함수
	go func() {
		for {
			tcpconn, _ := listener.Accept()
			buffer := make([]byte, 1024)
			for {
				n, err := tcpconn.Read(buffer)
				fmt.Println("test", string(buffer[:n]))
				if n == 0 {
					break
				}
				if err != nil {
					fmt.Println(err)
				}
				AChannel <- Message{Text: string(buffer[:n])}

			}
		}

	}()

	wg.Add(1)
	// b channel에 할당
	go func() {
		for item := range AChannel {
			time.Sleep(time.Second * 2)
			BChannel <- item
		}
	}()

	wg.Add(1)
	// udp 연결루틴 생성 및 송신
	go func() {
		defer wg.Done()
		for response := range BChannel {
			// 클라이언트로 msg write
			msg := []byte(response.Text)
			_, _ = udpconn.Write(msg)
			fmt.Println(string(msg))
		}
	}()
	wg.Wait()
}
