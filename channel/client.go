package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	conn, err := net.Dial("tcp", "localhost:3000")
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:3001")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Print("Enter a message: ")
	msg := readInput()
	_, _ = conn.Write([]byte(msg))

	udpconn, _ := net.DialUDP("udp", nil, serverAddr)

	defer udpconn.Close()
	fmt.Printf("The UDP server is %s\n", udpconn.RemoteAddr().String())
	wg.Add(1)

	go func() {
		defer wg.Done()

		buffer := make([]byte, 1024)
		for {
			// 파일을 읽을때 까지 대기함
			n, _, _ := udpconn.ReadFromUDP(buffer)
			message := string(buffer[:n])
			//_, _ = conn.Write([]byte(message))
			fmt.Println("Received message(UDP client)/ReadFrom:", message)

		}
	}()
	wg.Wait()

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
