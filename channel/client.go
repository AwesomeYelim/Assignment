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
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001")

	defer conn.Close()

	udpconn, _ := net.ListenUDP("udp", serverAddr)

	defer udpconn.Close()

	buffer := make([]byte, 1024)
	for {
		fmt.Print("Enter a message: ")
		msg := readInput()
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
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
