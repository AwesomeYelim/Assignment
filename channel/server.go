package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	channel()
}

type Message struct {
	Text string
}

var wg sync.WaitGroup

func Ahandler(conn net.Conn, wg *sync.WaitGroup, ch chan Message) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputStr := string(buffer[:n])
	fmt.Println("Received from client:", inputStr)

	ch <- Message{Text: inputStr}

	time.Sleep(time.Second)

	response := <-ch
	fmt.Println("Sending response to client:", response)
	wg.Add(1)
	go Bhandler(wg, ch)
	close(ch)
	//conn.Write([]byte(response.Text))
}

func Bhandler(wg *sync.WaitGroup, ch chan Message) {
	defer wg.Done()
	BChannel := make(chan Message, 2)
	response := <-ch
	BChannel <- response

	fmt.Println(response)

}

func channel() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error is", err)
		return
	}
	defer listener.Close()

	AChannel := make(chan Message, 2)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		wg.Add(1)
		go Ahandler(conn, &wg, AChannel)
		wg.Wait()
	}

}
