package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- strconv.Itoa(i)
	}
}
