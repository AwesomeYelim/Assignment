package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")

	conn, _ := net.DialUDP("udp", nil, serverAddr)

	defer conn.Close()

	for {
		fmt.Print("Enter a message: ")
		message := readInput()
		_, _ = conn.Write([]byte(message))

		//fmt.Println("Message sent to UDP server:", message)
		//buffer := make([]byte, 1024)
		//n, ServerAddr, _ := conn.ReadFromUDP(buffer)
		//receivedMessage := string(buffer[:n])
		//fmt.Printf("Received message from %s: %s\n", ServerAddr.String(), receivedMessage)
	}

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
