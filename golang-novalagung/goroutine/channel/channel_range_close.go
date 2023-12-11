package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

/*
channel direction
ch chan string = this param can use for send and receive data
ch chan <- string = this param only use for send data
ch <-chan string = this param only use for receive data
*/
