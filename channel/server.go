package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	channel()
}

type Message struct {
	Text string
}

func channel() {
	var wg sync.WaitGroup

	listener, _ := net.Listen("tcp", "localhost:3000")        // tcp 서버 생성
	udpconn, _ := net.ResolveUDPAddr("udp", "localhost:3001") // udp 생성

	defer listener.Close()

	// a,b channel 을 생성
	AChannel := make(chan Message, 2) // 초기 버퍼를 설정하는 이유는 => 처음 채널 크기는 0이므로 데이터를 빼갈때 까지 대기함 => 데이터를 가져가지 않아서 프로그램이 멈추는 현상이 생긴다(deadlock)
	BChannel := make(chan Message, 2)

	for {
		tcpconn, err := listener.Accept()
		udpconn, _ := net.ListenUDP("udp", udpconn)

		if err != nil {
			continue
		}

		wg.Add(1)
		// 통신을 받아서 a channel 에 할당하는 함수
		go func(conn net.Conn, wg *sync.WaitGroup) {
			defer conn.Close()

			buffer := make([]byte, 1024)

			n, _ := conn.Read(buffer)
			AChannel <- Message{Text: string(buffer[:n])}
			fmt.Println("A-channel 전달 완료^_^", <-AChannel)

			time.Sleep(time.Second)

			//conn.Write([]byte(response.Text))
		}(tcpconn, &wg)

		wg.Add(1)
		// b channel에 할당
		go func(wg *sync.WaitGroup) {
			defer tcpconn.Close()
			BChannel <- <-AChannel
		}(&wg)

		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("UDP server/ReadFrom ", udpconn.LocalAddr())
			buffer := make([]byte, 1024)
			for {
				n, clientAddr, _ := udpconn.ReadFrom(buffer)
				receivedMessage := string(buffer[:n])
				fmt.Printf("Received message from %s: %s\n", clientAddr.String(), receivedMessage)
			}
		}()

		wg.Add(1)
		// udp 연결루틴 생성 및 송신
		go func() {
			defer wg.Done()
			fmt.Println("UDP server/WriteTo", udpconn.LocalAddr())
			for {
				response := <-BChannel
				// 클라이언트로 msg write
				msg := []byte(response.Text)
				_, err = udpconn.WriteTo(msg, udpconn.LocalAddr())
				if err != nil {
					log.Fatal(err)
				}
			}
		}()

		wg.Wait()
	}

}
