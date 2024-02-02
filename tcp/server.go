package main

import (
	"fmt"
	"net"
	"os"
)

func handler(conn net.Conn) {
	buffer := make([]byte, 1024) // 데이터를 담을 buffer 생성

	for {
		n, err := conn.Read(buffer) // conn은 메모리 주소값을 참조한다. 루프를 도는 이유는 순차적으로  읽기 => 쓰기(보내기) 를 진행하기 위함
		if err != nil {
			defer os.Exit(0) // 인터럽트 신호로 터미널 닫음
			//fmt.Println(err)
			return
		}
		if n > 0 {
			fmt.Println("server : ", string(buffer[:n]))
			conn.Write(buffer[:n]) // buffer 형식 내 데이터 값이 있는 index 까지 잘라주어 client 로 보낸다. => byte 형식으로 보내야함
		}

	}

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
