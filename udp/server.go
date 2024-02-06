package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")

	conn, _ := net.ListenUDP("udp", serverAddr)

	defer conn.Close()

	fmt.Println("UDP server is listening on", serverAddr)

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buffer)
		receivedMessage := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", clientAddr.String(), receivedMessage)

		//fmt.Print("Enter a message: ")
		//message := readInput()
		//
		//_, _ = conn.Write([]byte(message))
		//
		//fmt.Println("Message sent to UDP server:", message)
	}
}
