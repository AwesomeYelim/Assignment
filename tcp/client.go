package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	localAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}

	// 서버에 연결
	conn, _ := net.DialTCP("tcp", localAddr, &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234})

	for {
		fmt.Print("Enter a message: ")
		msg := readInput()
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
