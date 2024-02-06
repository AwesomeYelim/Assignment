package main

import (
	"fmt"
	"net"
)

type Message struct {
	Text string
}

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	listener, _ := net.Listen("tcp", "localhost:3000")
	defer listener.Close()

	conn, _ := net.ListenUDP("udp", serverAddr)
	tcpconn, _ := listener.Accept()
	defer conn.Close()

	fmt.Println("UDP server is listening on", serverAddr)

	// a,b channel 을 생성
	AChannel := make(chan Message, 2) // 초기 버퍼를 설정하는 이유는 => 처음 채널 크기는 0이므로 데이터를 빼갈때 까지 대기함 => 데이터를 가져가지 않아서 프로그램이 멈추는 현상이 생긴다(deadlock)
	BChannel := make(chan Message, 2)

	go func() {
		recv := make([]byte, 1024)
		defer tcpconn.Close()
		for {
			n, _ := tcpconn.Read(recv)
			AChannel <- Message{string(recv[:n])}
		}
	}()

	go func() {
		for {
			BChannel <- <-AChannel
		}
	}()

	go func() {
		t := <-BChannel
		msg := []byte(t.Text)
		_, _ = conn.WriteTo(msg, serverAddr)
		fmt.Println(t)
	}()
	//fmt.Print("Enter a message: ")
	//message := readInput()
	//
	//_, _ = conn.Write([]byte(message))
	//
	//fmt.Println("Message sent to UDP server:", message)

}
