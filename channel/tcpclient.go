package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", "localhost:3000")
	conn, _ := net.DialTCP("tcp", localAddr, &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234})

	defer conn.Close()

	for {
		fmt.Print("Enter a message: ")
		msg := readInput()
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
