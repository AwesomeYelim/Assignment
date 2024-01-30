package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Read data from the client
	buffer := make([]byte, 1024) // 6. 1024 크기의 바이트 슬라이스 생성
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println(string(buffer[:n]))
	// 7. client 쪽으로 buffer를 쏴준다.
	_, err = conn.Write(buffer[:n])
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	// Close the connection
	conn.Close()
}

func main() {
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", "localhost:8080") // 1. TCP 프로토콜에 8080  포트로 연결을 받음
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close() // 2. main 함수가 끝나기 직전에 연결 대기를 닫음

	fmt.Println("Server listening on port 8080")

	for { // 무한 루프를 생성한다. 들어오는 연결을 지속적으로 수신해야하는 서버의 일반적인 패턴임
		conn, err := listener.Accept() // 3. client 가 연결되면 TCP 연결을 리턴
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		defer conn.Close() // 4. main 함수가 끝나기 직전에 TCP 연결을 닫음

		go handleConnection(conn) // 5. 패킷을 처리할 함수를 고루틴으로 실행
	}
}
