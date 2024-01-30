package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // 1
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	go func(c net.Conn) {
		send := []string{"피카츄", "라이츄", "파이리", "꼬부기"}
		for i := range send {
			_, err = c.Write([]byte(send[i]))
			if err != nil {
				fmt.Println("Failed to write data : ", err)
				break
			}
			if i == len(send)-1 {
				defer conn.Close()
				return
			}
			i++
			time.Sleep(5 * time.Second)

		}
	}(conn)

	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("client : ", string(recv[:n]))
		}
	}(conn)

	fmt.Scanln()
}
