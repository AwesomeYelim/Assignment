package main

import (
	"fmt"
	"net"
)

func handler(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n > 0 {
			fmt.Println("server : ", string(buffer[:n]))
			conn.Write(buffer[:n])
		}
	}

	//fmt.Println(string(buffer[:n]))
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer conn.Close()

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080") // 1
	if err != nil {
		fmt.Println("Error is", err)
		return
	}
	defer listener.Close()

	for { // 무한 루프를 생성한다. 들어오는 연결을 지속적으로 수신해야하는 서버의 일반적인 패턴임
		conn, err := listener.Accept() // 접속을 대기한다. => 사용자 접속시 handler 함수로 해당 커넥션을 처리한다.
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handler(conn)
	}

}
