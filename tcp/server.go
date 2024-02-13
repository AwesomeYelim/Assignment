package main

import (
	"fmt"
	"net"
)

func main() {
	listenAddr, err := net.ResolveTCPAddr("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error is", err)
		return
	}
	// 서버 시작
	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		buffer := make([]byte, 1024)

		for {
			fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
			n, err := conn.Read(buffer)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("server : ", string(buffer[:n]))

		}
	}
}
