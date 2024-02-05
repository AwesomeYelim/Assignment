package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Dial : 전화걸다

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	inputMsg := bufio.NewScanner(os.Stdin)
	inputMsg.Scan()
	input := inputMsg.Text()
	_, err = conn.Write([]byte(input))

	//serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	//if err != nil {
	//	fmt.Println("Error resolving UDP address:", err)
	//	return
	//}
	//
	//clientConn, err := net.DialUDP("udp", nil, serverAddr)
	//if err != nil {
	//	fmt.Println("Error connecting to UDP server:", err)
	//	return
	//}
	//defer clientConn.Close()
	//
	//buffer := make([]byte, 1024)
	//
	//n, _, err := clientConn.ReadFromUDP(buffer)
	//if err != nil {
	//	fmt.Println("err(UDP client):", err)
	//	return
	//}
	//message := string(buffer[:n])
	//fmt.Println("Received message(UDP client):", message)
}
