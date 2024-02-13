package main

import (
	"fmt"
	"net"
)

func main() {
	// 서버 주소 설정
	serverAddr, err := net.ResolveTCPAddr("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// 로컬 주소 설정
	localAddr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}

	// 서버에 연결
	conn, err := net.DialTCP("tcp", localAddr, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()
	fmt.Println(serverAddr, localAddr)
}
