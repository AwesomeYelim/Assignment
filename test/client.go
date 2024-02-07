package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:3000")
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001")

	defer conn.Close()

	udpconn, _ := net.ListenUDP("udp", serverAddr)

	defer udpconn.Close()

	buffer := make([]byte, 1024)
	for {
		fmt.Print("Enter a message: ")
		msg := readInput()
		_, _ = conn.Write([]byte(msg))
		// 파일을 읽을때 까지 대기함
		n, _, _ := udpconn.ReadFromUDP(buffer)
		message := string(buffer[:n])
		fmt.Println("Received message(UDP client)/ReadFrom:", message)
	}

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
