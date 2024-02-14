package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001")

	udpconn, _ := net.ListenUDP("udp", serverAddr)

	defer udpconn.Close()

	buffer := make([]byte, 1024)
	for {
		// 파일을 읽을때 까지 대기함
		n, _, _ := udpconn.ReadFromUDP(buffer)
		message := string(buffer[:n])
		fmt.Println("Received message(UDP client)/ReadFrom:", message)
	}

}
