package main

import "fmt"

func main() {
	channel()
}

func channel() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "안녕"
		myChannel <- "예림"

		close(myChannel)
	}()

	for msg := range myChannel {
		fmt.Println(msg)
	}
}
